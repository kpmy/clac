package clac

import (
	"math/big"
)

func init_MIXED_IC() {
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

	put(MULT, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Mul(lr, rc.Re)
		return
	}))))

	put(QUOT, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = rc.Im
		ret.Re = lr.Quo(lr, rc.Re)
		return
	}))))

	put(POW, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr *big.Rat, rc *Cmp) (ret *Cmp) {
		panic("wrong")
	}))))

	put(LSS, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(LEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(GEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(GTR, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(EQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))

	put(NEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr *big.Rat, rc *Cmp) bool {
		panic("wrong")
	}))))
}
