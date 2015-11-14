package clac

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
		panic("wrong")
	}))))

	put(LSS, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))

	put(LEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))

	put(GEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))

	put(GTR, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))

	put(EQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))

	put(NEQ, COMPLEX, INTEGER, b_(b_c_(b_c_ir_(func(lc complex128, rr float64) bool {
		panic("wrong")
	}))))
}
