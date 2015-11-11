package clac

import (
	"github.com/kpmy/ypk/halt"
	"math/big"
)

const (
	less = -1
	eq   = 0
	gtr  = 1
)

func b_(fn func(Value, Value) bool) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: BOOLEAN}
		ret.V = fn(l, r)
		return
	}
}

func b_i_(fn func(*big.Int, Value) bool) func(Value, Value) bool {
	return func(l Value, r Value) bool {
		li := l.ToInt()
		return fn(li, r)
	}
}

func b_i_i_(fn func(*big.Int, *big.Int) bool) func(*big.Int, Value) bool {
	return func(li *big.Int, r Value) bool {
		ri := r.ToInt()
		return fn(li, ri)
	}
}

func i_(fn func(Value, Value) *big.Int) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: INTEGER}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func i_i_(fn func(*big.Int, Value) *big.Int) func(Value, Value) *big.Int {
	return func(l Value, r Value) *big.Int {
		li := l.ToInt()
		return fn(li, r)
	}
}

func i_i_i_(fn func(*big.Int, *big.Int) *big.Int) func(*big.Int, Value) *big.Int {
	return func(li *big.Int, r Value) *big.Int {
		ri := r.ToInt()
		return fn(li, ri)
	}
}

func r_(fn func(Value, Value) *big.Rat) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: RATIONAL}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func r_r_(fn func(*big.Rat, Value) *big.Rat) func(Value, Value) *big.Rat {
	return func(l Value, r Value) *big.Rat {
		lr := l.ToRat()
		return fn(lr, r)
	}
}

func r_r_r_(fn func(*big.Rat, *big.Rat) *big.Rat) func(*big.Rat, Value) *big.Rat {
	return func(lr *big.Rat, r Value) *big.Rat {
		rr := r.ToRat()
		return fn(lr, rr)
	}
}

func r_ir_(fn func(*big.Rat, Value) *big.Rat) func(Value, Value) *big.Rat {
	return func(l Value, r Value) *big.Rat {
		if l.T == RATIONAL {
			lr := l.ToRat()
			return fn(lr, r)
		} else if l.T == INTEGER {
			li := l.ToInt()
			lr := big.NewRat(0, 1)
			lr.SetInt(li)
			return fn(lr, r)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func r_ir_ir_(fn func(*big.Rat, *big.Rat) *big.Rat) func(*big.Rat, Value) *big.Rat {
	return func(lr *big.Rat, r Value) *big.Rat {
		if r.T == RATIONAL {
			rr := r.ToRat()
			return fn(lr, rr)
		} else if r.T == INTEGER {
			ri := r.ToInt()
			rr := big.NewRat(0, 1)
			rr.SetInt(ri)
			return fn(lr, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_(fn func(Value, Value) *Cmp) func(Value, Value) Value {
	return func(l Value, r Value) (ret Value) {
		ret = Value{T: COMPLEX}
		ret.V = thisVal(fn(l, r))
		return
	}
}

func c_c_(fn func(*Cmp, Value) *Cmp) func(Value, Value) *Cmp {
	return func(l Value, r Value) *Cmp {
		lc := l.ToCmp()
		return fn(lc, r)
	}
}

func c_c_c_(fn func(*Cmp, *Cmp) *Cmp) func(*Cmp, Value) *Cmp {
	return func(lc *Cmp, r Value) *Cmp {
		rc := r.ToCmp()
		return fn(lc, rc)
	}
}

func c_ir_(fn func(*big.Rat, Value) *Cmp) func(Value, Value) *Cmp {
	return func(l Value, r Value) *Cmp {
		if l.T == RATIONAL {
			lr := l.ToRat()
			return fn(lr, r)
		} else if l.T == INTEGER {
			li := l.ToInt()
			lr := big.NewRat(0, 1)
			lr.SetInt(li)
			return fn(lr, r)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_ir_ir_(fn func(*big.Rat, *big.Rat) *Cmp) func(*big.Rat, Value) *Cmp {
	return func(lr *big.Rat, r Value) *Cmp {
		if r.T == RATIONAL {
			rr := r.ToRat()
			return fn(lr, rr)
		} else if r.T == INTEGER {
			ri := r.ToInt()
			rr := big.NewRat(0, 1)
			rr.SetInt(ri)
			return fn(lr, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_c_ir_(fn func(*Cmp, *big.Rat) *Cmp) func(*Cmp, Value) *Cmp {
	return func(lc *Cmp, r Value) *Cmp {
		if r.T == RATIONAL {
			rr := r.ToRat()
			return fn(lc, rr)
		} else if r.T == INTEGER {
			ri := r.ToInt()
			rr := big.NewRat(0, 1)
			rr.SetInt(ri)
			return fn(lc, rr)
		} else {
			halt.As(100, "cannot convert")
		}
		panic(0)
	}
}

func c_ir_c_(fn func(*big.Rat, *Cmp) *Cmp) func(*big.Rat, Value) *Cmp {
	return func(lr *big.Rat, r Value) *Cmp {
		rc := r.ToCmp()
		return fn(lr, rc)
	}
}
