package unit

import "math/big"

func f64(n float64) *big.Rat {
	return new(big.Rat).SetFloat64(n)
}

var lengthUnits = []baseUnit{
	{
		name:      "m",
		aliases:   []string{"meter", "metre"},
		value:     f64(1),
		dimension: LENGTH,
		prefixes:  &metricPrefixes,
	},
	{
		name:      "in",
		aliases:   []string{"inch", "inches"},
		value:     f64(0.0254),
		dimension: LENGTH,
	},
	{
		name:      "ft",
		aliases:   []string{"foot", "feet"},
		value:     f64(0.3048),
		dimension: LENGTH,
	},
	{
		name:      "yd",
		aliases:   []string{"yard", "yards"},
		value:     f64(0.9144),
		dimension: LENGTH,
	},
	{
		name:      "mi",
		aliases:   []string{"mile", "miles"},
		value:     f64(1609.344),
		dimension: LENGTH,
	},
}
