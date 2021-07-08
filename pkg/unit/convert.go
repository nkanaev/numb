package unit

import (
	"fmt"
	"math/big"
)

func ratexp(x *big.Rat, n int) *big.Rat {
	if n == 0 {
		return new(big.Rat).SetInt64(1)
	}
	r := new(big.Rat).Set(x)
	for i := 1; i < n; i++ {
		r.Mul(r, x)
	}
	if n < 0 {
		r.Quo(new(big.Rat).SetInt64(1), r)
	}
	return r
}

func Convert(n *big.Rat, from, to UnitList) *big.Rat {
	if !from.Conforms(to) {
		panic(fmt.Sprintf("incompatible units: %s & %s", from, to))
	}
	if from.String() == to.String() {
		return n
	}
	return to.denormalize(from.normalize(n))
}
