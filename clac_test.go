package clac

import (
	"testing"
)

func TestClac(t *testing.T) {
	i := This(INTEGER, 4)
	t.Log(Do(NEG, i))
	r := This(RATIONAL, 5.8)
	t.Log(Do(NEG, r))
}
