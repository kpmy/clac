package clac

import (
	"math/big"
)

func init_MIXED_RI() {
	put(SUM, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))

	put(MULT, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Mul(lr, rr)
	}))))

	put(QUOT, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Quo(lr, rr)
	}))))

	put(POW, RATIONAL, INTEGER, r_(r_ir_(r_ir_ir_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		panic("wrong")
	}))))

	put(LSS, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == lss
	}))))

	put(LEQ, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != gtr
	}))))

	put(GEQ, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != lss
	}))))

	put(GTR, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == gtr
	}))))

	put(EQ, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == eq
	}))))

	put(NEQ, RATIONAL, INTEGER, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != eq
	}))))
}
