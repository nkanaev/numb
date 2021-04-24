package unit

import (
	"fmt"
	"math/big"
)

type Dimension uint

const (
	UNKNOWN Dimension = 1 << iota
	LENGTH
	TEMPERATURE
	VOLUME
	MASS
	TIME
	ANGLE
	DIGITAL
)

type Unit struct {
	name string
	value, offset float64
	dimension Dimension
}

func (u *Unit) String() string {
	return u.name
}

var meter = &Unit{name: "m", value: 1, offset: 0, dimension: LENGTH}
var inch = &Unit{name: "in", value: 0.0254, offset: 0, dimension: LENGTH}

func Get(x string) *Unit {
	if x == "m" {
		return meter
	}
	if x == "in" {
		return inch
	}
	return nil
}

func normalize(n *big.Rat, u *Unit) *big.Rat {
	// (n + u.offset) * u.value
	offset := new(big.Rat).SetFloat64(u.offset)
	value := new(big.Rat).SetFloat64(u.value)
	num := new(big.Rat).Set(n)
	num = num.Add(num, offset)
	return num.Mul(num, value)
}

func denormalize(n *big.Rat, u *Unit) *big.Rat {
	// (n / u.value) - u.offset
	offset := new(big.Rat).SetFloat64(u.offset)
	value := new(big.Rat).SetFloat64(u.value)
	num := new(big.Rat).Set(n)
	num = num.Quo(num, value)
	return num.Sub(num, offset)
}

func Convert(n *big.Rat, from, to *Unit) *big.Rat {
	if from.dimension != to.dimension {
		panic(fmt.Sprintf("incompatible units: %s & %s", from, to))
	}
	if from.name == to.name {
		return new(big.Rat).Set(n)
	}
	// normalize
	return denormalize(normalize(n, from), to)
}
