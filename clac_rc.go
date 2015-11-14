package clac

func init_MIXED_RC() {
	put(SUM, FLOAT, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) + rc
	}))))

	put(DIFF, FLOAT, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) - rc
	}))))

	put(MULT, FLOAT, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) * rc
	}))))

	put(QUOT, FLOAT, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		return complex(lr, 0) / rc
	}))))

	put(POW, FLOAT, COMPLEX, c_(c_ir_(c_ir_c_(func(lr float64, rc complex128) (ret complex128) {
		panic("wrong")
	}))))

	put(LSS, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(LEQ, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(GEQ, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(GTR, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(EQ, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))

	put(NEQ, FLOAT, COMPLEX, b_(b_ir_(b_ir_c_(func(lr float64, rc complex128) bool {
		panic("wrong")
	}))))
}
