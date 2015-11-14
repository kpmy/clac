package clac

import (
	"math/cmplx"
)

func init_MIXED_IC() {
	put(SUM, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) + rc
	}))))

	put(DIFF, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) - rc
	}))))

	put(MULT, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) * rc
	}))))

	put(QUOT, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) / rc
	}))))

	put(POW, INTEGER, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return cmplx.Pow(complex(lr, 0), rc)
	}))))

	put(LSS, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(LEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(GEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(GTR, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(EQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(NEQ, INTEGER, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))
}
