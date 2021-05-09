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
		name:         "in²",
		shortaliases: []string{"in2", "sqin"},
		value:        f64(0.00064516),
		dimension:    AREA,
	},
	{
		name:         "ft²",
		shortaliases: []string{"ft2", "sqft"},
		aliases:      []string{"sqfeet"},
		value:        f64(0.09290304),
		dimension:    AREA,
	},
	{
		name:         "yd²",
		shortaliases: []string{"sqyd"},
		aliases:      []string{"sqyard"},
		value:        f64(0.83612736),
		dimension:    AREA,
	},
	{
		name:         "rd²",
		shortaliases: []string{"sqrd"},
		value:        f64(25.29295),
		dimension:    AREA,
	},
	{
		name:         "ch²",
		shortaliases: []string{"sqch"},
		value:        f64(404.6873),
		dimension:    AREA,
	},
	{
		name:         "mi²",
		shortaliases: []string{"sqmil", "sqmi"},
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
