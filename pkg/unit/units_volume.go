package unit

var volumeUnits = []baseUnit{
	{
		name:         "m³",
		shortaliases: []string{"m3"},
		value:        f64(1),
		dimension:    VOLUME,
		prefixes:     &metricPrefixes,
		prefixpow:    3,
	},
	{
		name:      "l",
		aliases:   []string{"liter", "litre"},
		value:     f64(0.001),
		dimension: VOLUME,
		prefixes:  &metricPrefixes,
	},
}
