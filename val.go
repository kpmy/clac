package clac

import (
	"github.com/kpmy/ypk/halt"
	"math/big"
	"reflect"
)

type Cmp struct {
	Re, Im *big.Rat
}

func (v Value) ToInt() (ret *big.Int) {
	switch v.T {
	case INTEGER:
		ret = big.NewInt(0)
		ret.Add(ret, v.V.(*big.Int))
	default:
		halt.As(100, v.T)
	}
	return
}

func (v Value) ToRat() (ret *big.Rat) {
	switch v.T {
	case RATIONAL:
		ret = big.NewRat(0, 1)
		ret.Add(ret, v.V.(*big.Rat))
	default:
		halt.As(100, v.T)

	}
	return
}

func (v Value) ToCmp() (ret *Cmp) {
	switch v.T {
	case COMPLEX:
		ret = new(Cmp)
		ret.Re = big.NewRat(0, 1)
		ret.Im = big.NewRat(0, 1)
		ret.Re.Add(ret.Re, v.V.(*Cmp).Re)
		ret.Im.Add(ret.Im, v.V.(*Cmp).Im)
	default:
		halt.As(100, v.T)

	}
	return
}

func This(typ Type, _x interface{}) (ret Value) {
	switch typ {
	case INTEGER:
		switch x := _x.(type) {
		case int:
			ret.V = big.NewInt(int64(x))
		default:
			halt.As(101, reflect.TypeOf(x))
		}
	case RATIONAL:
		switch x := _x.(type) {
		case float64:
			r := big.NewRat(0, 1)
			ret.V = r.SetFloat64(x)
		default:
			halt.As(102, reflect.TypeOf(x))
		}
	default:
		halt.As(100, typ)
	}
	ret.T = typ
	return
}

func thisVal(x interface{}) (ret interface{}) {
	switch x.(type) {
	case *big.Int, *big.Rat:
		ret = x
	default:
		halt.As(100, reflect.TypeOf(x))
	}
	return
}
