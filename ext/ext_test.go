package ext

import (
	"github.com/kpmy/clac"
	"testing"
)

const BYTE clac.Type = 100

const ASH clac.Op = 101

func init() {
	clac.Grow(ASH, BYTE, clac.INTEGER, func(l clac.Value, r clac.Value) (ret clac.Value) {
		ret.T = BYTE
		x := l.V.(uint8)
		n := r.V.(int64)
		if n > 0 {
			ret.V = uint8(x << uint(n))
		} else if n < 0 {
			ret.V = uint8(x >> uint(-n))
		} else {
			ret.V = x
		}
		return
	})
}

func Byte(x uint8) (ret clac.Value) {
	ret.T = BYTE
	ret.V = x
	return
}

func TestExtended(t *testing.T) {
	t.Log(clac.Do2(Byte(32), ASH, clac.Int(1)))
	t.Log(clac.Do2(Byte(32), ASH, clac.Int(-1)))
}
