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

func isDerived(u UnitList) bool {
	if len(u) == 0 {
		return false
	}
	return len(u) > 1 || u[0].Exp != 1
}

func normalize(n *big.Rat, units UnitList) *big.Rat {
	num := new(big.Rat).Set(n)
	if isDerived(units) {
		for _, u := range units {
			num.Mul(num, ratexp(u.Unit.value, u.Exp))
		}
	} else {
		// (n + u.offset) * u.value
		u := units[0].Unit
		if u.offset != nil {
			num.Add(num, u.offset)
		}
		return num.Mul(num, u.value)
	}
	return num
}

func denormalize(n *big.Rat, units UnitList) *big.Rat {
	num := new(big.Rat).Set(n)
	if isDerived(units) {
		for _, u := range units {
			num.Quo(num, ratexp(u.Unit.value, u.Exp))
		}
	} else {
		// (n / u.value) - u.offset
		u := units[0].Unit
		num.Quo(num, u.value)
		if u.offset != nil {
			num.Sub(num, u.offset)
		}
	}
	return num
}

func Convert(n *big.Rat, from, to *UnitList) *big.Rat {
	if !from.Conforms(to) {
		panic(fmt.Sprintf("incompatible units: %s & %s", from, to))
	}
	if from.String() == to.String() {
		return n
	}
	return denormalize(normalize(n, *from), *to)
}
