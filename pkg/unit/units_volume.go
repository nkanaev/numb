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
		name:         "l",
		shortaliases: []string{"lt"},
		aliases:      []string{"liter", "litre"},
		value:        f64(0.001),
		dimension:    VOLUME,
		prefixes:     &metricPrefixes,
	},
	{
		name:         "in³",
		shortaliases: []string{"in3", "cuin"},
		value:        f64(1.6387064e-5),
		dimension:    VOLUME,
	},
	{
		name:         "ft³",
		shortaliases: []string{"ft3", "cuft"},
		value:        f64(0.028316846592),
		dimension:    VOLUME,
	},
	{
		name:         "yd³",
		shortaliases: []string{"yd3", "cuyd"},
		value:        f64(0.764554857984),
		dimension:    VOLUME,
	},
	{
		name:      "teaspoon",
		value:     f64(0.000005),
		dimension: VOLUME,
	},
	{
		name:      "tablespoon",
		value:     f64(0.000015),
		dimension: VOLUME,
	},
}
