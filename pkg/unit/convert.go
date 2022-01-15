package unit

import (
	"fmt"
	"math/big"
)

type ConformanceError struct {
	a, b Units
}

func (c ConformanceError) Error() string {
	return fmt.Sprintf(
		"incompatible units, %s (%s) does not conform %s (%s)",
		c.a.String(), c.a.Dimension().Measure(),
		c.b.String(), c.b.Dimension().Measure())
}

func Convert(n *big.Rat, from, to Units) (*big.Rat, error) {
	if from.Dimension().IsPure() {
		return from.normalize(n), nil
	}
	if !from.Conforms(to) {
		return nil, ConformanceError{from, to}
	}
	if from.String() == to.String() {
		return n, nil
	}
	return to.denormalize(from.normalize(n)), nil
}

func Normalize(n *big.Rat, x Units) *big.Rat {
	return x.normalize(n)
}
