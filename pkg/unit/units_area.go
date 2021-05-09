package unit

var areaUnits = []baseUnit{
	{
		name:         "m²",
		shortaliases: []string{"m2"},
		value:        f64(1),
		dimension:    AREA,
		prefixes:     &metricPrefixes,
		prefixpow:    2,
	},
	{
		name:      "sqin",
		value:     f64(0.00064516),
		dimension: AREA,
	},
	{
		name:      "sqft",
		aliases:   []string{"sqfeet"},
		value:     f64(0.09290304),
		dimension: AREA,
	},
	{
		name:      "sqyd",
		aliases:   []string{"sqyard"},
		value:     f64(0.83612736),
		dimension: AREA,
	},
	{
		name:      "sqrd",
		value:     f64(25.29295),
		dimension: AREA,
	},
	{
		name:      "sqch",
		value:     f64(404.6873),
		dimension: AREA,
	},
	{
		name:         "sqmi",
		shortaliases: []string{"sqmil"},
		aliases:      []string{"sqmile"},
		value:        f64(6.4516e-10),
		dimension:    AREA,
	},
	{
		name:      "acre",
		value:     f64(4046.86),
		dimension: AREA,
	},
	{
		name:      "hectare",
		value:     f64(10000),
		dimension: AREA,
	},
}
