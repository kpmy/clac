package clac

func init_INTEGER() {
	put(NEG, INTEGER, NoType, i_(i_i_(func(li int64, r Value) int64 {
		return -li
	})))

	put(SUM, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li int64, ri int64) int64 {
		return li + ri
	}))))

	put(DIFF, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li int64, ri int64) int64 {
		return li - ri
	}))))

	put(MULT, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li int64, ri int64) int64 {
		return li * ri
	}))))

	put(DIV, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li int64, ri int64) int64 {
		return li / ri
	}))))

	put(MOD, INTEGER, INTEGER, i_(i_i_(i_i_i_(func(li int64, ri int64) int64 {
		return li % ri
	}))))

	put(EQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li == ri
	}))))

	put(NEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li != ri
	}))))

	put(GTR, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li > ri
	}))))

	put(LSS, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li < ri
	}))))

	put(GEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li >= ri
	}))))

	put(LEQ, INTEGER, INTEGER, b_(b_i_(b_i_i_(func(li int64, ri int64) bool {
		return li <= ri
	}))))
}
