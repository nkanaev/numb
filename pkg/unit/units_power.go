package unit

var powerUnits = []baseUnit{
	{
		name:        "W",
		aliases:     []string{"watt"},
		value:       f64(1),
		dimension:   POWER,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		name:        "hp",
		aliases:     []string{"horsepower"},
		value:       f64(745.69987158227),
		dimension:   POWER,
		description: "mechanical horsepower",
	},
}
