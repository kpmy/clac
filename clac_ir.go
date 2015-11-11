package clac

import (
	"math/big"
)

func init_MIXED_IR() {
	put(SUM, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))

	put(MULT, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Mul(lr, rr)
	}))))

	put(QUOT, INTEGER, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Quo(lr, rr)
	}))))

	put(QUOT, INTEGER, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Quo(lr, rr)
	}))))

	put(POW, INTEGER, RATIONAL, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		panic("wrong")
	}))))

	put(LSS, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == lss
	}))))

	put(LEQ, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != gtr
	}))))

	put(GEQ, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != lss
	}))))

	put(GTR, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == gtr
	}))))

	put(EQ, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == eq
	}))))

	put(NEQ, INTEGER, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != eq
	}))))
}
