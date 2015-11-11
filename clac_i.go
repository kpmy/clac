package clac

import (
	"math/big"
)

func init_INTEGER() {
	put(NEG, INTEGER, NoType, i_(i_i_(func(li *big.Int, r Value) *big.Int {
		return li.Neg(li)
	})))

	put(SUM, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Add(li, ri)
	}))))

	put(DIFF, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Sub(li, ri)
	}))))

	put(MULT, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Mul(li, ri)
	}))))

	put(DIV, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Div(li, ri)
	}))))

	put(MOD, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		return li.Mod(li, ri)

	}))))

	put(POW, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li *big.Int, ri *big.Int) *big.Int {
		panic("wrong")
	}))))

	put(EQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res == eq
	}))))

	put(NEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res != eq
	}))))

	put(GTR, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res == gtr
	}))))

	put(LSS, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res == lss
	}))))

	put(GEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res != lss
	}))))

	put(LEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li *big.Int, ri *big.Int) bool {
		res := li.Cmp(ri)
		return res != gtr
	}))))
}
