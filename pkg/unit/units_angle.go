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
		name:        "rad",
		aliases:     []string{"radian"},
		value:       f64(1),
		dimension:   ANGLE,
		description: "SI derived unit",
	},
	{
		name:        "°",
		aliases:     []string{"deg", "degree"},
		value:       unitdiv(consts.PI, 180),
		dimension:   ANGLE,
		description: "SI-accepted unit",
	},
	{
		name:        "arcsec",
		value:       unitdiv(consts.PI, 648000),
		dimension:   ANGLE,
		description: "SI-accepted unit (pi / 648000)",
	},
	{
		name:        "arcmin",
		value:       unitdiv(consts.PI, 10800),
		dimension:   ANGLE,
		description: "SI-accepted unit (pi / 10800)",
	},
	{
		name:      "grad",
		aliases: []string{"grade", "gradian"},
		value:     unitdiv(consts.PI, 200),
		dimension: ANGLE,
	},
	{
		name:      "cycle",
		value:     unitdiv(consts.PI, 2),
		dimension: ANGLE,
	},
}
