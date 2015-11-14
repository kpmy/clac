package clac

import (
	"math"
)

func init_MIXED_IR() {
	put(SUM, INTEGER, FLOAT, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr + rr
	}))))

	put(DIFF, INTEGER, FLOAT, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr - rr
	}))))

	put(MULT, INTEGER, FLOAT, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr * rr
	}))))

	put(QUOT, INTEGER, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return lr / rr
	}))))

	put(QUOT, INTEGER, FLOAT, r_(r_r_(r_r_r_(func(lr float64, rr float64) float64 {
		return lr / rr
	}))))

	put(POW, INTEGER, FLOAT, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return math.Pow(lr, rr)
	}))))

	put(POW, INTEGER, INTEGER, r_(r_ir_(r_ir_ir_(func(lr float64, rr float64) float64 {
		return math.Pow(lr, rr)
	}))))

	put(LSS, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr < rr
	}))))

	put(LEQ, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr <= rr
	}))))

	put(GEQ, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr >= rr
	}))))

	put(GTR, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr > rr
	}))))

	put(EQ, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr == rr
	}))))

	put(NEQ, INTEGER, FLOAT, b_(b_ir_(b_ir_ir_(func(lr float64, rr float64) bool {
		return lr != rr
	}))))
}
