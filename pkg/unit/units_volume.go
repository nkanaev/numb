package unit

var volumeUnits = []baseUnit{
	{
		short:     "m³, m3",
		value:     f64(1),
		dimension: VOLUME,
		prefixes:  &metricPrefixes,
		prefixpow: 3,
	},
	{
		short:       "l, lt",
		long:        "liter, litre",
		value:       f64(0.001),
		dimension:   VOLUME,
		prefixes:    &metricPrefixes,
		description: "SI-accepted unit (1 l = 0.001 m³)",
	},
	{
		short:     "in³, in3, cuin",
		value:     f64(1.6387064e-5),
		dimension: VOLUME,
	},
	{
		short:     "ft³, ft3, cuft",
		value:     f64(0.028316846592),
		dimension: VOLUME,
	},
	{
		short:     "yd³, yd3, cuyd",
		value:     f64(0.764554857984),
		dimension: VOLUME,
	},
	{
		short:     "teaspoon",
		value:     f64(0.000005),
		dimension: VOLUME,
	},
	{
		short:     "tablespoon",
		value:     f64(0.000015),
		dimension: VOLUME,
	},
}
