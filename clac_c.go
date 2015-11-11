package clac

func init_COMPLEX() {
	put(NEG, COMPLEX, NoType, c_(c_c_(func(lc *Cmp, r Value) *Cmp {
		ret := new(Cmp)
		ret.Im = lc.Im.Neg(lc.Im)
		ret.Re = lc.Re.Neg(lc.Re)
		return ret
	})))

	put(CON, COMPLEX, NoType, c_(c_c_(func(lc *Cmp, r Value) *Cmp {
		ret := new(Cmp)
		ret.Im = ret.Im.Neg(lc.Im)
		ret.Re = lc.Re
		return ret
	})))

	put(SUM, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		ret := new(Cmp)
		ret.Re = lc.Re.Add(lc.Re, rc.Re)
		ret.Im = lc.Im.Add(lc.Im, rc.Im)
		return ret
	}))))

	put(DIFF, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		ret := new(Cmp)
		ret.Re = lc.Re.Sub(lc.Re, rc.Re)
		ret.Im = lc.Im.Sub(lc.Im, rc.Im)
		return ret
	}))))

	put(MULT, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		panic("wrong")
		ret := new(Cmp)
		ret.Re = lc.Re.Mul(lc.Re, rc.Re)
		ret.Im = lc.Im.Mul(lc.Im, rc.Im)
		return ret
	}))))

	put(QUOT, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		panic("wrong")
		ret := new(Cmp)
		ret.Re = lc.Re.Quo(lc.Re, rc.Re)
		ret.Im = lc.Im.Quo(lc.Im, rc.Im)
		return ret
	}))))

	put(POW, COMPLEX, COMPLEX, c_(c_c_(c_c_c_(func(rc *Cmp, lc *Cmp) *Cmp {
		panic("wrong")
	}))))

	put(LSS, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))

	put(LEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))

	put(GEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))

	put(GTR, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))

	put(EQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))

	put(NEQ, COMPLEX, COMPLEX, b_(b_c_(b_c_c_(func(rc *Cmp, lc *Cmp) bool {
		panic("wrong")
	}))))
}
