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
		return cmplx.Pow(lc, rc)
	}))))

	put(EQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc complex128, lc complex128) bool {
		return rc == lc
	}))))

	put(NEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(lc complex128, rc complex128) bool {
		return lc != rc
	}))))
}
