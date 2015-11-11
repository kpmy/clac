package clac

import (
	"math/big"
)

func init_RATIONAL() {
	put(NEG, RATIONAL, NoType, r_(r_r_(func(lr *big.Rat, r Value) *big.Rat {
		return lr.Neg(lr)
	})))

	put(SUM, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Add(lr, rr)
	}))))

	put(DIFF, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Sub(lr, rr)
	}))))

	put(MULT, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Mul(lr, rr)
	}))))

	put(QUOT, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		return lr.Quo(lr, rr)
	}))))

	put(POW, RATIONAL, RATIONAL, r_(r_r_(r_r_r_(func(lr *big.Rat, rr *big.Rat) *big.Rat {
		panic("wrong")
	}))))

	put(LSS, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == lss
	}))))

	put(LEQ, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != gtr
	}))))

	put(GEQ, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != lss
	}))))

	put(GTR, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == gtr
	}))))

	put(EQ, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res == eq
	}))))

	put(NEQ, RATIONAL, RATIONAL, b_(b_ir_(b_ir_ir_(func(lr *big.Rat, rr *big.Rat) bool {
		res := lr.Cmp(rr)
		return res != eq
	}))))
}
