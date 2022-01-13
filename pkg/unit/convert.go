package unit

import (
	"fmt"
	"math/big"
)

type ConformanceError struct {
	a, b UnitList
}

func (c ConformanceError) Error() string {
	dim1, _ := c.a.Dimension().Measure()
	dim2, _ := c.b.Dimension().Measure()

	return fmt.Sprintf(
		"%s (%s) does not conform %s (%s)",
		c.a.String(), dim1.String(),
		c.b.String(), dim2.String())
}

func Convert(n *big.Rat, from, to UnitList) (*big.Rat, error) {
	if from.Dimension().IsZero() {
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

func Normalize(n *big.Rat, x UnitList) *big.Rat {
	return x.normalize(n)
}
