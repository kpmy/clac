package clac

import (
	"fmt"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"log"
	"strconv"
	"strings"
)

type Op int

type Type int

type Value struct {
	T Type
	V interface{}
}

func (v Value) String() string {
	switch v.T {
	case INTEGER:
		return fmt.Sprint(v.V)
	case FLOAT:
		return fmt.Sprint(v.V)
	case COMPLEX:
		return fmt.Sprint(v.V)
	case BOOLEAN:
		return fmt.Sprint(v.V)
	case ErrType:
		return "error"
	default:
		return fmt.Sprint("(", v.T, ")", v.V)
	}
}

type key struct {
	op Op
	l  Type
	r  Type
}

const ErrType Type = -1

const (
	NoType Type = iota
	INTEGER
	FLOAT
	COMPLEX

	BOOLEAN
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
	case FLOAT:
		return "FLOAT"
	case BOOLEAN:
		return "BOOLEAN"
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

type dfs struct {
	f   dfn
	err bool
}

type dfn func(Value, Value) Value

var om map[key]dfs
var dummy Value

func put2(op Op, l Type, r Type, fn dfn, err bool) {
	key := key{op, l, r}
	if _, ok := om[key]; !ok {
		om[key] = dfs{f: fn, err: err}
	} else {
		halt.As(100, "op already exists ", key)
	}
}

func put(op Op, l Type, r Type, fn dfn) {
	key := key{op, l, r}
	if _, ok := om[key]; !ok {
		om[key] = dfs{f: fn, err: false}
	} else {
		halt.As(100, "op already exists ", key)
	}
}

func Int(x int64) Value {
	return This(INTEGER, x)
}

func Flo(x float64) Value {
	return This(FLOAT, x)
}

func Cpx(x complex128) Value {
	return This(COMPLEX, x)
}

func Grow(op Op, l Type, r Type, fn dfn) {
	put(op, l, r, fn)
}

func init_ERR() {

	err := func(op Op, r Type, l Type) {
		put2(op, r, l, func(Value, Value) Value {
			halt.As(100, "ЕГГОГ", " ", op, " ", l, " ", r)
			panic(0)
		}, true)
	}

	err(DIV, FLOAT, FLOAT)
	err(DIV, FLOAT, INTEGER)
	err(DIV, INTEGER, FLOAT)
	err(DIV, COMPLEX, COMPLEX)
	err(DIV, FLOAT, COMPLEX)
	err(DIV, COMPLEX, FLOAT)
	err(DIV, COMPLEX, INTEGER)
	err(DIV, INTEGER, COMPLEX)

	err(MOD, FLOAT, FLOAT)
	err(MOD, FLOAT, INTEGER)
	err(MOD, INTEGER, FLOAT)
	err(MOD, COMPLEX, COMPLEX)
	err(MOD, FLOAT, COMPLEX)
	err(MOD, COMPLEX, FLOAT)
	err(MOD, COMPLEX, INTEGER)
	err(MOD, INTEGER, COMPLEX)

	for op := NoOp + 1; op < lastOp; op++ {
		if op.Monadic() {
			err(op, BOOLEAN, NoType)
		} else {
			for t := NoType + 1; t < lastType; t++ {
				err(op, BOOLEAN, t)
				if t != BOOLEAN {
					err(op, t, BOOLEAN)
				}
			}
		}
	}

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
	err(CON, FLOAT, NoType)

	for op := LSS; op <= GTR; op++ {
		for t := NoType + 1; t < lastType; t++ {
			if t != BOOLEAN && op != EQ && op != NEQ {
				err(op, t, COMPLEX)
				if t != COMPLEX {
					err(op, COMPLEX, t)
				}
			}
		}
	}
}

func init() {
	om = make(map[key]dfs)
	init_ERR()

	init_INTEGER()
	init_FLOAT()
	init_COMPLEX()

	init_MIXED_IR()
	init_MIXED_IC()

	init_MIXED_RI()
	init_MIXED_RC()

	init_MIXED_CI()
	init_MIXED_CR()

	log.SetFlags(0)

	for i := NoOp + 1; i < lastOp; i++ {
		for l := NoType + 1; l < lastType; l++ {
			for r := NoType; r < lastType; r++ {
				if r != NoType || i.Monadic() {
					k := key{i, l, r}
					if _, ok := om[k]; !ok {
						panic(k)
						log.Println(k)
					}
				}
			}
		}
	}
}

func wrap(f dfn, left, right Value) (ret Value) {
	defer func() {
		if x := recover(); x != nil {
			if strings.Contains(fmt.Sprint(x), "divide by zero") {
				//ok
			} else {
				panic(x)
			}
		}
	}()

	ret.T = ErrType
	ret = f(left, right)
	return
}

func Do(op Op, value Value) (ret Value) {
	assert.For(op.Monadic(), 20)
	//log.Println(op, value)
	if f, ok := om[key{op: op, l: value.T}]; ok {
		ret = wrap(f.f, value, dummy)
	} else {
		halt.As(100, op, value)
	}
	//log.Println(ret)
	return
}

func Do2(left Value, op Op, right Value) (ret Value) {
	assert.For(!op.Monadic(), 20)
	//log.Println(left, op, right)
	if f, ok := om[key{op: op, l: left.T, r: right.T}]; ok {
		ret = wrap(f.f, left, right)
	} else {
		halt.As(100, op, left, right)
	}
	//log.Println(ret)
	return
}

func Forbidden(op Op, typ ...Type) bool {
	switch len(typ) {
	case 0:
		return true
	case 1:
		if !op.Monadic() {
			return true
		} else {
			if f, ok := om[key{op: op, l: typ[0]}]; ok {
				return f.err
			}
		}
	case 2:
		if op.Monadic() {
			return Forbidden(op, typ[0])
		} else {
			if f, ok := om[key{op: op, l: typ[0], r: typ[1]}]; ok {
				return f.err
			}
		}
	}
	return false
}
