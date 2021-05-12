package unit

var energyUnits = []baseUnit{
	{
		name:      "J",
		aliases:   []string{"joule"},
		value:     f64(1),
		dimension: ENERGY,
		prefixes:  &metricPrefixes,
		description: "SI derived unit",
	},
	{
		name:      "Wh",
		aliases:   []string{"watt-hour"},
		value:     f64(3600),
		dimension: ENERGY,
		prefixes:  &metricPrefixes,
	},
	{
		name:      "eV",
		aliases:   []string{"electronvolt"},
		value:     f64(1.602176565e-19),
		dimension: ENERGY,
		prefixes:  &metricPrefixes,
	},
	{
		name:      "erg",
		value:     exp(10, -7),
		dimension: ENERGY,
	},
}
