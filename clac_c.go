package clac

import (
	"math/cmplx"
)

func init_COMPLEX() {
	put(NEG, COMPLEX, NoType, c_(c_c_(func(lc complex128, r Value) complex128 {
		return -lc
	})))

	put(CON, COMPLEX, NoType, c_(c_c_(func(lc complex128, r Value) complex128 {
		return cmplx.Conj(lc)
	})))

	put(SUM, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(lc complex128, rc complex128) complex128 {
		return lc + rc
	}))))

	put(DIFF, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(lc complex128, rc complex128) complex128 {
		return lc - rc
	}))))

	put(MULT, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(lc complex128, rc complex128) complex128 {
		return lc * rc
	}))))

	put(QUOT, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(lc complex128, rc complex128) complex128 {
		return lc / rc
	}))))

	put(POW, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(lc complex128, rc complex128) complex128 {
		panic("wrong")
	}))))

	put(LSS, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		panic("wrong")
	}))))

	put(LEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		panic("wrong")
	}))))

	put(GEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		panic("wrong")
	}))))

	put(GTR, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		panic("wrong")
	}))))

	put(EQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc complex128, lc complex128) bool {
		panic("wrong")
	}))))

	put(NEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		panic("wrong")
	}))))
}
