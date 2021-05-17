package unit

import "math/big"

func f64(n float64) *big.Rat {
	return new(big.Rat).SetFloat64(n)
}

var lengthUnits = []baseUnit{
	{
		short:       "m",
		long:        "meter, metre",
		value:       f64(1),
		dimension:   LENGTH,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
	{
		short:     "in",
		long:      "inch",
		value:     f64(0.0254),
		dimension: LENGTH,
	},
	{
		short:     "ft",
		long:      "foot, feet",
		value:     f64(0.3048),
		dimension: LENGTH,
	},
	{
		short:     "yd",
		long:      "yard, yards",
		value:     f64(0.9144),
		dimension: LENGTH,
	},
	{
		short:     "mi",
		long:      "mile, miles",
		value:     f64(1609.344),
		dimension: LENGTH,
	},

	{
		short:     "li",
		long:      "link",
		value:     f64(0.201168),
		dimension: LENGTH,
	},
	{
		short:     "rd",
		long:      "rod",
		value:     f64(5.0292),
		dimension: LENGTH,
	},
	{
		short:     "ch",
		long:      "chain",
		value:     f64(20.1),
		dimension: LENGTH,
	},
	{
		short:     "angstrom",
		value:     exp(10, -10),
		dimension: LENGTH,
	},
	{
		short:     "mil",
		value:     f64(0.0000254),
		dimension: LENGTH,
	},
	{
		short:       "au",
		long:        "astronomical-unit",
		value:       f64(149597870700),
		dimension:   LENGTH,
		description: "SI-accepted unit",
	},
	{
		short:       "ly",
		long:        "lightyeaar, light-year",
		value:       f64(9460730472580800),
		dimension:   LENGTH,
		prefixes:    &metricPrefixes,
		description: "SI-accepted unit",
	},
	{
		short:     "lightsecond",
		long:      "lightsecond, light-second",
		value:     f64(299792458),
		dimension: LENGTH,
		prefixes:  &metricPrefixes,
	},
}
