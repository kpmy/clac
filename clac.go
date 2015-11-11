package clac

import (
	"fmt"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"log"
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
	case RATIONAL:
		return "RATIONAL"
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
	err(CON, RATIONAL, NoType)
}

func init() {
	om = make(map[key]dfn)
	init_ERR()

	init_INTEGER()
	init_RATIONAL()
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

func Do(op Op, value Value) (ret Value) {
	assert.For(op.Monadic(), 20)
	log.Println(op, value)
	if f, ok := om[key{op: op, l: value.T}]; ok {
		ret = f(value, dummy)
	} else {
		halt.As(100, op, value)
	}
	log.Println(ret)
	return
}

func Do2(left Value, op Op, right Value) (ret Value) {
	assert.For(!op.Monadic(), 20)
	log.Println(left, op, right)
	if f, ok := om[key{op: op, l: left.T, r: right.T}]; ok {
		ret = f(left, right)
	} else {
		halt.As(100, op, left, right)
	}
	log.Println(ret)
	return
}
