package clac

import (
	"math/big"
)

func init_MIXED_RC() {
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

	put(MULT, RATIONAL, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Mul(lr, rc.Re)
		return
	}))))

	put(QUOT, RATIONAL, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Quo(lr, rc.Re)
		return
	}))))

	put(POW, RATIONAL, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
	}))))

	put(LSS, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(LEQ, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(GEQ, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(GTR, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(EQ, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(NEQ, RATIONAL, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))
}
