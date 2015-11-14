package clac

import (
	"github.com/kpmy/ypk/halt"
	"reflect"
)

func (v Value) ToBool() (ret bool) {
	switch v.T {
	case BOOLEAN:
		ret = v.V.(bool)
	default:
		halt.As(100, v.T)
	}
	return
}

func (v Value) ToInt() (ret int64) {
	switch v.T {
	case INTEGER:
		ret = v.V.(int64)
	default:
		halt.As(100, v.T)
	}
	return
}

func (v Value) ToFlo() (ret float64) {
	switch v.T {
	case INTEGER:
		ret = float64(v.V.(int64))
	case FLOAT:
		ret = v.V.(float64)
	default:
		halt.As(100, v.T)

	}
	return
}

func (v Value) ToCmp() (ret complex128) {
	switch v.T {
	case COMPLEX:
		ret = v.V.(complex128)
	default:
		halt.As(100, v.T)

	}
	return
}

func This(typ Type, _x interface{}) (ret Value) {
	switch typ {
	case INTEGER:
		switch x := _x.(type) {
		case int64:
			ret.V = x
		case int:
			ret.V = int64(x)
		default:
			halt.As(101, reflect.TypeOf(x))
		}
	case FLOAT:
		switch x := _x.(type) {
		case float64:
			ret.V = x
		default:
			halt.As(102, reflect.TypeOf(x))
		}
	case COMPLEX:
		switch x := _x.(type) {
		case complex128:
			ret.V = x
		}
	default:
		halt.As(100, typ)
	}
	ret.T = typ
	return
}

func thisVal(x interface{}) (ret interface{}) {
	switch x.(type) {
	case int64, complex128, float64:
		ret = x
	default:
		halt.As(100, reflect.TypeOf(x))
	}
	return
}
