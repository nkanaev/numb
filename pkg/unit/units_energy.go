package unit

var energyUnits = []baseUnit{
	{
		short:       "J",
		long:        "joule",
		value:       f64(1),
		dimension:   ENERGY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		short:     "Wh",
		long:      "watt-hour",
		value:     f64(3600),
		dimension: ENERGY,
		prefixes:  &metricPrefixes,
	},
	{
		short:       "eV",
		long:        "electronvolt",
		value:       f64(1.602176565e-19),
		dimension:   ENERGY,
		prefixes:    &metricPrefixes,
		description: "SI-accepted unit",
	},
	{
		short:     "erg",
		value:     exp(10, -7),
		dimension: ENERGY,
	},
}
