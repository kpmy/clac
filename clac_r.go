package clac

func init_FLOAT() {
	put(NEG, FLOAT, NoType, r_(r_r_(func(lr float64, r Value) float64 {
		return -lr
	})))

	put(SUM, FLOAT, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		return lr + rr
	}))))

	put(DIFF, FLOAT, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		return lr - rr
	}))))

	put(MULT, FLOAT, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		return lr * rr
	}))))

	put(QUOT, FLOAT, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		return lr / rr
	}))))

	put(POW, FLOAT, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		panic("wrong")
	}))))

	put(LSS, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr < rr
	}))))

	put(LEQ, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr <= rr
	}))))

	put(GEQ, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr >= rr
	}))))

	put(GTR, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr > rr
	}))))

	put(EQ, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr == rr
	}))))

	put(NEQ, FLOAT, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr != rr
	}))))
}
