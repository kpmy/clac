package clac

import (
	"github.com/kpmy/ypk/assert"
	"testing"
)

func TestClac(t *testing.T) {
	i := This(INTEGER, 4)
	t.Log(Do(NEG, i))
	r := This(RATIONAL, 5.8)
	t.Log(Do(NEG, r))
}

func Int(x int) Value {
	return This(INTEGER, x)
}

func Rat(x float64) Value {
	return This(RATIONAL, x)
}

func Cpx(x complex128) Value {
	return This(COMPLEX, x)
}

func TestAny(t *testing.T) {
	assert.For(Do2(Do2(Int(4), SUM, Int(24)), EQ, Int(4+24)).ToBool(), 20)
	assert.For(Do2(Do2(Int(4), DIFF, Int(24)), EQ, Int(4-24)).ToBool(), 20)
	assert.For(Do2(Do2(Int(4), MULT, Int(24)), EQ, Int(4*24)).ToBool(), 20)
	t.Log(Do2(Int(4), QUOT, Int(24)))

	t.Log(Rat(34.5))
	t.Log(Cpx(complex(3, 5)))
}
