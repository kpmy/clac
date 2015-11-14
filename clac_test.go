package clac

import (
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"math/rand"
	"testing"
	"time"
)

func TestClac(t *testing.T) {
	i := This(INTEGER, 4)
	t.Log(Do(NEG, i))
	r := This(FLOAT, 5.8)
	t.Log(Do(NEG, r))
}

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Rnd(t Type) (ret Value) {
	switch t {
	case INTEGER:
		ret = Int(rnd.Int63n(10) - 5)
	case FLOAT:
		ret = Flo(10*rnd.Float64() - 5)
	case COMPLEX:
		ret = Cpx(complex(10*rnd.Float64()-5, 10*rnd.Float64()-5))
	default:
		halt.As(100, t)
	}
	return
}

func TestSome(t *testing.T) {
	assert.For(Do2(Do2(Int(4), SUM, Int(24)), EQ, Int(4+24)).ToBool(), 20)
	assert.For(Do2(Do2(Int(4), DIFF, Int(24)), EQ, Int(4-24)).ToBool(), 20)
	assert.For(Do2(Do2(Int(4), MULT, Int(24)), EQ, Int(4*24)).ToBool(), 20)
	t.Log(Do2(Int(4), QUOT, Int(24)))
	t.Log(Do(NEG, Int(45)))
	t.Log(Flo(34.5))
	t.Log(Cpx(complex(3, 5)))
	t.Log(Do(CON, Cpx(complex(3.5, 14))))
}

func TestAny(t *testing.T) {
	for l := NoType + 1; l < lastType; l++ {
		for r := NoType + 1; r < lastType; r++ {
			for op := NoOp + 1; op < lastOp; op++ {
				if l != BOOLEAN && r != BOOLEAN {
					if !Forbidden(op, l, r) {
						if op.Monadic() {
							if l == r {
								rnd := Rnd(l)
								t.Log(op, l, rnd, "=", Do(op, rnd))
							}
						} else {
							lv := Rnd(l)
							rv := Rnd(r)
							t.Log(l, lv, op, r, rv, "=", Do2(lv, op, rv))
						}
					}
				}
			}
		}
	}
}
