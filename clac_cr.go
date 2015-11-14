package clac

import (
	"math/cmplx"
)

func init_MIXED_CR() {
	put(SUM, COMPLEX, FLOAT, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc + complex(rr, 0)
	}))))

	put(DIFF, COMPLEX, FLOAT, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc - complex(rr, 0)
	}))))

	put(MULT, COMPLEX, FLOAT, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc * complex(rr, 0)
	}))))

	put(QUOT, COMPLEX, FLOAT, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc / complex(rr, 0)
	}))))

	put(POW, COMPLEX, FLOAT, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return cmplx.Pow(lc, complex(rr, 0))
	}))))

	put(EQ, COMPLEX, FLOAT, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		return complex(rr, 0) == lc
	}))))

	put(NEQ, COMPLEX, FLOAT, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		return complex(rr, 0) == lc
	}))))
}
