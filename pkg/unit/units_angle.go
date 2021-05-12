package unit

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/consts"
)

func unitdiv(a *big.Rat, x int64) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, big.NewRat(x, 1))
	return num
}

var angleUnits = []baseUnit{
	{
		name:      "rad",
		aliases:   []string{"radian"},
		value:     f64(1),
		dimension: ANGLE,
		description: "SI derived unit",
	},
	{
		name:      "deg",
		aliases:   []string{"degree"},
		value:     unitdiv(consts.PI, 180),
		dimension: ANGLE,
	},
}
