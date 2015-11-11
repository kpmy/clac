package clac

import (
	"math/big"
)

func init_MIXED_CR() {
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

	put(MULT, COMPLEX, RATIONAL, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Mul(rr, lc.Re)
		return
	}))))

	put(QUOT, COMPLEX, RATIONAL, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
		ret = new(Cmp)
		ret.Im = lc.Im
		ret.Re = rr.Quo(lc.Re, rr)
		return
	}))))

	put(POW, COMPLEX, RATIONAL, c_(c_c_(c_c_ir_(func(lc *Cmp, rr *big.Rat) (ret *Cmp) {
		panic("wrong")
	}))))

	put(LSS, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(LEQ, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(GEQ, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(GTR, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(EQ, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))

	put(NEQ, COMPLEX, RATIONAL, b_(b_c_(b_c_ir_(func(lc *Cmp, rr *big.Rat) bool {
		panic("wrong")
	}))))
}
