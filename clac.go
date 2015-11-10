package clac

import (
	"fmt"
	"github.com/kpmy/ypk/halt"
	"log"
	"math/big"
	"strconv"
)

type Op int

type Type int

type Value struct {
	T Type
	V interface{}
}

type key struct {
	op Op
	l  Type
	r  Type
}

const (
	NoType Type = iota
	INTEGER
	RATIONAL
	COMPLEX

	lastType
)

const (
	NoOp Op = iota

	NEG
	CON

	SUM
	DIFF
	MULT
	QUOT
	DIV
	MOD
	POW

	LSS
	LEQ
	EQ
	NEQ
	GEQ
	GTR

	lastOp
)

func (k key) String() string {
	return fmt.Sprint(k.l, k.op, k.r)
}

func (t Type) String() string {
	switch t {
	case NoType:
		return "NONE"
	case INTEGER:
		return "INTEGER"
	case COMPLEX:
		return "COMPLEX"
	case RATIONAL:
		return "RATIONAL"
	default:
		return fmt.Sprint("type ", strconv.Itoa(int(t)))
	}
}

func (o Op) String() string {
	switch o {
	case SUM:
		return "+"
	case DIFF:
		return "-"
	case MULT:
		return "*"
	case QUOT:
		return "/"
	case DIV:
		return "//"
	case MOD:
		return "%"
	case POW:
		return "^"
	case NEG:
		return "--"
	case CON:
		return "_"
	case GTR:
		return ">"
	case GEQ:
		return ">="
	case EQ:
		return "="
	case NEQ:
		return "#"
	case LEQ:
		return "<="
	case LSS:
		return "<"
	default:
		return fmt.Sprint("op ", strconv.Itoa(int(o)))
	}
}

func (o Op) Monadic() bool {
	return o == NEG || o == CON
}

type dfn func(Value, Value) Value

var om map[key]dfn
var dummy Value

func put(op Op, l Type, r Type, fn dfn) {
	key := key{op, l, r}
	if _, ok := om[key]; !ok {
		om[key] = fn
	} else {
		halt.As(100, "op already exists ", key)
	}
}

func init_INTEGER() {
	put(NEG, INTEGER, NoType, i_(i_i_(func(li *big.Int, r Value) *big.Int {
		return li.Neg(li)
	})))

	put(SUM, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Add(li, ri)
	}))))

	put(DIFF, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Sub(li, ri)
	}))))
}

func init_RATIONAL() {
	put(NEG, RATIONAL, NoType, r_(r_r_(func(lr *big.Rat, r Value) *big.Rat {
		return lr.Neg(lr)
	})))

	put(SUM, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))
}

func init_COMPLEX() {
	put(NEG, COMPLEX, NoType, c_(c_c_(func(lc *Cmp, r Value) *Cmp {
		ret := new(Cmp)
		ret.Im = lc.Im.Neg(lc.Im)
		ret.Re = lc.Re.Neg(lc.Re)
		return ret
	})))

	put(CON, COMPLEX, NoType, c_(c_c_(func(lc *Cmp, r Value) *Cmp {
		ret := new(Cmp)
		ret.Im = ret.Im.Neg(lc.Im)
		ret.Re = lc.Re
		return ret
	})))

	put(SUM, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		ret := new(Cmp)
		ret.Re = lc.Re.Add(lc.Re, rc.Re)
		ret.Im = lc.Im.Add(lc.Im, rc.Im)
		return ret
	}))))

	put(DIFF, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		ret := new(Cmp)
		ret.Re = lc.Re.Sub(lc.Re, rc.Re)
		ret.Im = lc.Im.Sub(lc.Im, rc.Im)
		return ret
	}))))
}

func init_MIXED() {
	put(SUM, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))

	put(SUM, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Add(lr, rc.Re)
		return
	}))))

	put(DIFF, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Sub(lr, rc.Re)
		return
	}))))

	put(SUM, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))

	put(SUM, RATIONAL, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Add(lr, rc.Re)
		return
	}))))

	put(DIFF, RATIONAL, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Sub(lr, rc.Re)
		return
	}))))

	put(SUM, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Add(rr, lc.Re)
		return
	}))))

	put(DIFF, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Sub(lc.Re, rr)
		return
	}))))

	put(SUM, COMPLEX, RATIONAL, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Add(rr, lc.Re)
		return
	}))))

	put(DIFF, COMPLEX, RATIONAL, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Sub(lc.Re, rr)
		return
	}))))
}

func init_ERR() {
	_e := func(Value, Value) Value {
		halt.As(100, "ЕГГОГ")
		panic(0)
	}

	err := func(op Op, r Type, l Type) {
		put(op, r, l, _e)
	}

	err(DIV, RATIONAL, RATIONAL)
	err(DIV, RATIONAL, INTEGER)
	err(DIV, INTEGER, RATIONAL)
	err(DIV, COMPLEX, COMPLEX)
	err(DIV, RATIONAL, COMPLEX)
	err(DIV, COMPLEX, RATIONAL)
	err(DIV, COMPLEX, INTEGER)
	err(DIV, INTEGER, COMPLEX)

	err(MOD, RATIONAL, RATIONAL)
	err(MOD, RATIONAL, INTEGER)
	err(MOD, INTEGER, RATIONAL)
	err(MOD, COMPLEX, COMPLEX)
	err(MOD, RATIONAL, COMPLEX)
	err(MOD, COMPLEX, RATIONAL)
	err(MOD, COMPLEX, INTEGER)
	err(MOD, INTEGER, COMPLEX)

	for op := NoOp + 1; op < lastOp; op++ {
		if op.Monadic() {
			for l := NoType + 1; l < lastType; l++ {
				for r := NoType + 1; r < lastType; r++ {
					err(op, l, r)
				}
			}
		}
	}

	for l := NoType; l < lastType; l++ {
		for r := NoType; r < lastType; r++ {
			err(NoOp, l, r)
		}
	}

	for o := NoOp + 1; o < lastOp; o++ {
		if o.Monadic() {
			for r := NoType + 1; r < lastType; r++ {
				err(o, NoType, r)
			}
		} else {
			for t := NoType + 1; t < lastType; t++ {
				err(o, NoType, t)
				err(o, t, NoType)
			}
		}
	}

	err(CON, INTEGER, NoType)
	err(CON, RATIONAL, NoType)
}

func init() {
	om = make(map[key]dfn)
	init_INTEGER()
	init_RATIONAL()
	init_COMPLEX()
	init_MIXED()

	init_ERR()

	log.SetFlags(0)
	for i := NoOp + 1; i < lastOp; i++ {
		for l := NoType + 1; l < lastType; l++ {
			for r := NoType; r < lastType; r++ {
				if r != NoType || i.Monadic() {
					k := key{i, l, r}
					if _, ok := om[k]; !ok {
						panic(k)
					}
				}
			}
		}
	}
}

func Do(op Op, value Value) (ret Value) {
	if f, ok := om[key{op: op, l: value.T}]; ok {
		ret = f(value, dummy)
	} else {
		halt.As(100, op, value)
	}
	return
}
