package unit

var radioactivityUnits = []baseUnit{
	{
		name:        "Bq",
		aliases:     []string{"becquierel"},
		value:       f64(1),
		dimension:   RADIOACTIVITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		name:    "Ci",
		aliases: []string{"curie"},
		value:   f64(3.7e+10),
		dimension:   RADIOACTIVITY,
	},
	{
		name:    "Rd",
		aliases: []string{"rutherford"},
		value:   exp(10, 6),
		dimension:   RADIOACTIVITY,
	},
}
