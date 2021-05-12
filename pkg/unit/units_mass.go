package unit

import "math/big"

func parse(x string) *big.Rat {
	rat, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse: " + x)
	}
	return rat
}

var massUnits = []baseUnit{
	{
		name:        "g",
		aliases:     []string{"gram"},
		value:       f64(0.001),
		dimension:   MASS,
		prefixes:    &metricPrefixes,
		description: "(0.001 kg - SI base unit)",
	},
	{
		name:      "t",
		aliases:   []string{"tonne", "metric-ton"},
		value:     f64(1000),
		dimension: MASS,
		prefixes:  &metricPrefixes,
		description: "SI-mentioned unit (1 t = 1000 kg)",
	},
	{
		name: "Da",
		shortaliases: []string{"u"},
		aliases: []string{"dalton"},
		value: parse("1.6605402e-27"),
		dimension: MASS,
	},
}
