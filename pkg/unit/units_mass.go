package unit

var massUnits = []baseUnit{
	{
		name:      "g",
		aliases:   []string{"gram"},
		value:     f64(0.001),
		dimension: MASS,
		prefixes:  &metricPrefixes,
	},
}
