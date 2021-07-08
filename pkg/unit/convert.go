package unit

import (
	"fmt"
	"math/big"
)

func Convert(n *big.Rat, from, to UnitList) *big.Rat {
	if !from.Conforms(to) {
		panic(fmt.Sprintf("incompatible units: %s & %s", from, to))
	}
	if from.String() == to.String() {
		return n
	}
	return to.denormalize(from.normalize(n))
}
