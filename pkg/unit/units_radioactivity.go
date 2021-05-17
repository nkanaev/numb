package unit

var radioactivityUnits = []baseUnit{
	{
		short:       "Bq",
		long:        "becquierel",
		value:       f64(1),
		dimension:   RADIOACTIVITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		short:     "Ci",
		long:      "curie",
		value:     f64(3.7e+10),
		dimension: RADIOACTIVITY,
	},
	{
		short:     "Rd",
		long:      "rutherford",
		value:     exp(10, 6),
		dimension: RADIOACTIVITY,
	},
}
