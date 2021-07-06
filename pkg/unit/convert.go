package unit

import (
	//"fmt"
	"math/big"
)

func normalize(n *big.Rat, u *NamedUnit) *big.Rat {
	// (n + u.offset) * u.value
	num := new(big.Rat).Set(n)
	if u.offset != nil {
		num.Add(num, u.offset)
	}
	return num.Mul(num, u.value)
}

func denormalize(n *big.Rat, u *NamedUnit) *big.Rat {
	// (n / u.value) - u.offset
	num := new(big.Rat).Set(n)
	num.Quo(num, u.value)
	if u.offset != nil {
		num.Sub(num, u.offset)
	}
	return num
}

func Convert(n *big.Rat, from, to *Unit) *big.Rat {
	return n
	/*
	if from.dimension != to.dimension {
		panic(fmt.Sprintf("incompatible units: %s & %s", from, to))
	}
	if from.name == to.name {
		return n
	}
	// normalize
	return denormalize(normalize(n, from), to)
	*/
}
