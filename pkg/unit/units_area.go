package unit

var areaUnits = []baseUnit{
	{
		name:      "m²",
		shortaliases:   []string{"m2"},
		value:     f64(1),
		dimension: AREA,
		prefixes:  &metricPrefixes,
		prefixpow: 2,
	},
}
