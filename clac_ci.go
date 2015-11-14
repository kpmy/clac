package clac

import (
	"math/cmplx"
)

func init_MIXED_CI() {
	put(SUM, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc + complex(rr, 0)
	}))))

	put(DIFF, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc - complex(rr, 0)
	}))))

	put(MULT, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc * complex(rr, 0)
	}))))

	put(QUOT, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return lc / complex(rr, 0)
	}))))

	put(POW, COMPLEX, INTEGER, c_(c_c_(c_c_ir_(func(lc complex128, rr float64) (ret complex128) {
		return cmplx.Pow(lc, complex(rr, 0))
	}))))

	put(EQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		return complex(rr, 0) == lc
	}))))

	put(NEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		return complex(rr, 0) == lc
	}))))
}
