package unit

var massUnits = []baseUnit{
	{
		name:      "g",
		aliases:   []string{"gram"},
		value:     f64(0.001),
		dimension: MASS,
		prefixes:  &metricPrefixes,
		description: "(0.001 kg - SI base unit)",
	},
	{
		name:      "t",
		aliases:   []string{"tonne", "metric-ton"},
		value:     f64(1000),
		dimension: MASS,
		prefixes:  &metricPrefixes,
	},
}
