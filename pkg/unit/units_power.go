package unit

var powerUnits = []baseUnit{
	{
		short:       "W",
		long:        "watt",
		value:       f64(1),
		dimension:   POWER,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		short:       "hp",
		long:        "horsepower",
		value:       f64(745.69987158227),
		dimension:   POWER,
		description: "mechanical horsepower",
	},
}
