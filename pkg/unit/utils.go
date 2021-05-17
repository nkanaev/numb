package unit

import "math/big"

var one = big.NewRat(1, 1)

func unitdiv(a *big.Rat, x int64) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, big.NewRat(x, 1))
	return num
}

func parse(x string) *big.Rat {
	rat, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse: " + x)
	}
	return rat
}

func f64(n float64) *big.Rat {
	return new(big.Rat).SetFloat64(n)
}
