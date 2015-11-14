package clac

import (
	"github.com/kpmy/ypk/halt"
)

func b_(fn func(Value, Value) bool) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: BOOLEAN}
		ret.V = fn(l, r)
		return
	}
}

func b_i_(fn func(int64, Value) bool) func(Value, Value) bool {
	return func(l Value, r Value) bool {
		li := l.ToInt()
		return fn(li, r)
	}
}

func b_i_i_(fn func(int64, int64) bool) func(int64, Value) bool {
	return func(li int64, r Value) bool {
		ri := r.ToInt()
		return fn(li, ri)
	}
}

func b_r_(fn func(float64, Value) bool) func(Value, Value) bool {
	return func(l Value, r Value) bool {
		lr := l.ToFlo()
		return fn(lr, r)
	}
}

func b_r_r_(fn func(float64, float64) bool) func(float64, Value) bool {
	return func(lr float64, r Value) bool {
		rr := r.ToFlo()
		return fn(lr, rr)
	}
}

func b_ir_(fn func(float64, Value) bool) func(Value, Value) bool {
	return func(l Value, r Value) bool {
		if l.T == FLOAT {
			lr := l.ToFlo()
			return fn(lr, r)
		} else if l.T == INTEGER {
			lr := l.ToFlo()
			return fn(lr, r)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func b_ir_ir_(fn func(float64, float64) bool) func(float64, Value) bool {
	return func(lr float64, r Value) bool {
		if r.T == FLOAT {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else if r.T == INTEGER {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func b_c_(fn func(complex128, Value) bool) func(Value, Value) bool {
	return func(l Value, r Value) bool {
		lc := l.ToCmp()
		return fn(lc, r)
	}
}

func b_c_c_(fn func(complex128, complex128) bool) func(complex128, Value) bool {
	return func(lc complex128, r Value) bool {
		rc := r.ToCmp()
		return fn(lc, rc)
	}
}

func b_c_ir_(fn func(complex128, float64) bool) func(complex128, Value) bool {
	return func(lc complex128, r Value) bool {
		if r.T == FLOAT {
			rr := r.ToFlo()
			return fn(lc, rr)
		} else if r.T == INTEGER {
			rr := r.ToFlo()
			return fn(lc, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func b_ir_c_(fn func(float64, complex128) bool) func(float64, Value) bool {
	return func(lr float64, r Value) bool {
		rc := r.ToCmp()
		return fn(lr, rc)
	}
}

func i_(fn func(Value, Value) int64) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: INTEGER}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func i_i_(fn func(int64, Value) int64) func(Value, Value) int64 {
	return func(l Value, r Value) int64 {
		li := l.ToInt()
		return fn(li, r)
	}
}

func i_i_i_(fn func(int64, int64) int64) func(int64, Value) int64 {
	return func(li int64, r Value) int64 {
		ri := r.ToInt()
		return fn(li, ri)
	}
}

func r_(fn func(Value, Value) float64) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: FLOAT}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func r_r_(fn func(float64, Value) float64) func(Value, Value) float64 {
	return func(l Value, r Value) float64 {
		lr := l.ToFlo()
		return fn(lr, r)
	}
}

func r_r_r_(fn func(float64, float64) float64) func(float64, Value) float64 {
	return func(lr float64, r Value) float64 {
		rr := r.ToFlo()
		return fn(lr, rr)
	}
}

func r_ir_(fn func(float64, Value) float64) func(Value, Value) float64 {
	return func(l Value, r Value) float64 {
		if l.T == FLOAT {
			lr := l.ToFlo()
			return fn(lr, r)
		} else if l.T == INTEGER {
			lr := l.ToFlo()
			return fn(lr, r)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func r_ir_ir_(fn func(float64, float64) float64) func(float64, Value) float64 {
	return func(lr float64, r Value) float64 {
		if r.T == FLOAT {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else if r.T == INTEGER {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_(fn func(Value, Value) complex128) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: COMPLEX}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func c_c_(fn func(complex128, Value) complex128) func(Value, Value) complex128 {
	return func(l Value, r Value) complex128 {
		lc := l.ToCmp()
		return fn(lc, r)
	}
}

func c_c_c_(fn func(complex128, complex128) complex128) func(complex128, Value) complex128 {
	return func(lc complex128, r Value) complex128 {
		rc := r.ToCmp()
		return fn(lc, rc)
	}
}

func c_ir_(fn func(float64, Value) complex128) func(Value, Value) complex128 {
	return func(l Value, r Value) complex128 {
		if l.T == FLOAT {
			lr := l.ToFlo()
			return fn(lr, r)
		} else if l.T == INTEGER {
			lr := l.ToFlo()
			return fn(lr, r)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_ir_ir_(fn func(float64, float64) complex128) func(float64, Value) complex128 {
	return func(lr float64, r Value) complex128 {
		if r.T == FLOAT {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else if r.T == INTEGER {
			rr := r.ToFlo()
			return fn(lr, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_c_ir_(fn func(complex128, float64) complex128) func(complex128, Value) complex128 {
	return func(lc complex128, r Value) complex128 {
		if r.T == FLOAT {
			rr := r.ToFlo()
			return fn(lc, rr)
		} else if r.T == INTEGER {
			rr := r.ToFlo()
			return fn(lc, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_ir_c_(fn func(float64, complex128) complex128) func(float64, Value) complex128 {
	return func(lr float64, r Value) complex128 {
		rc := r.ToCmp()
		return fn(lr, rc)
	}
}
