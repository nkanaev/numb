package unit

import "math/big"

func f64(n float64) *big.Rat {
	return new(big.Rat).SetFloat64(n)
}

var lengthUnits = []baseUnit{
	{
		name: "m",
		aliases: []string{"meter", "metre"},
		value: f64(1),
		dimension: LENGTH,
		prefixes: &metricPrefixes,
	},
	{
		name: "in",
		aliases: []string{"inch"},
		value: f64(0.0254),
		dimension: LENGTH,
	},
}
