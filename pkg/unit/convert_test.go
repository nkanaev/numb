package unit

import (
	"math/big"
	"testing"
)

func TestConvert(t *testing.T) {
	have := Convert(big.NewRat(1, 1), Get("km"), Get("m"))
	want := big.NewRat(1000, 1)
	if have.Cmp(want) != 0 {
		t.Errorf("\nwant: %s\nhave: %s", want, have)
	}
}
