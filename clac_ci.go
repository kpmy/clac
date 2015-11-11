package clac

import (
	"math/big"
)

func init_MIXED_CI() {
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

	put(MULT, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Mul(rr, lc.Re)
		return
	}))))

	put(QUOT, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Quo(lc.Re, rr)
		return
	}))))

	put(POW, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
	}))))

	put(LSS, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(LEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(GEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(GTR, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(EQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(NEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))
}
