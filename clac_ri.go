package clac

func init_MIXED_RI() {
	put(SUM, FLOAT, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr + rr
	}))))

	put(DIFF, FLOAT, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr - rr
	}))))

	put(MULT, FLOAT, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr * rr
	}))))

	put(QUOT, FLOAT, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr / rr
	}))))

	put(POW, FLOAT, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		panic("wrong")
	}))))

	put(LSS, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr < rr
	}))))

	put(LEQ, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr <= rr
	}))))

	put(GEQ, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr >= rr
	}))))

	put(GTR, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr > rr
	}))))

	put(EQ, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr == rr
	}))))

	put(NEQ, FLOAT, INTEGER, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr != rr
	}))))
}
