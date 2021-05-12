package unit

import "math/big"

func f64(n float64) *big.Rat {
	return new(big.Rat).SetFloat64(n)
}

var lengthUnits = []baseUnit{
	{
		name:        "m",
		aliases:     []string{"meter", "metre"},
		value:       f64(1),
		dimension:   LENGTH,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
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

	{
		name:      "li",
		aliases:   []string{"link"},
		value:     f64(0.201168),
		dimension: LENGTH,
	},
	{
		name:      "rd",
		aliases:   []string{"rod"},
		value:     f64(5.0292),
		dimension: LENGTH,
	},
	{
		name:      "ch",
		aliases:   []string{"chain"},
		value:     f64(20.1),
		dimension: LENGTH,
	},
	{
		name:      "angstrom",
		value:     exp(10, -10),
		dimension: LENGTH,
	},
	{
		name:      "mil",
		value:     f64(0.0000254),
		dimension: LENGTH,
	},
	{
		name:        "au",
		aliases:     []string{"astronomical-unit"},
		value:       f64(149597870700),
		dimension:   LENGTH,
		description: "SI-accepted unit",
	},
	{
		name:        "ly",
		aliases:     []string{"light-year"},
		value:       f64(9460730472580800),
		dimension:   LENGTH,
		prefixes:    &metricPrefixes,
		description: "SI-accepted unit",
	},
	{
		name:        "lightsecond",
		aliases:     []string{"lightsecond", "light-second"},
		value:       f64(299792458),
		dimension:   LENGTH,
		prefixes:    &metricPrefixes,
	},
}
