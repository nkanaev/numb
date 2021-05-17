package unit

var areaUnits = []baseUnit{
	{
		short:     "m², m2",
		value:     f64(1),
		dimension: AREA,
		prefixes:  &metricPrefixes,
		prefixpow: 2,
	},
	{
		short:     "in², in2, sqin",
		value:     f64(0.00064516),
		dimension: AREA,
	},
	{
		short:     "ft², ft2, sqft",
		long:      "sqfeet",
		value:     f64(0.09290304),
		dimension: AREA,
	},
	{
		short:     "yd², sqyd",
		long:      "sqyard",
		value:     f64(0.83612736),
		dimension: AREA,
	},
	{
		short:     "rd², rd2, sqrd",
		value:     f64(25.29295),
		dimension: AREA,
	},
	{
		short:     "ch², sqch",
		value:     f64(404.6873),
		dimension: AREA,
	},
	{
		short:     "mi², sqmil, sqmi",
		long:      "sqmile",
		value:     f64(6.4516e-10),
		dimension: AREA,
	},
	{
		short:     "acre",
		value:     f64(4046.86),
		dimension: AREA,
	},
	{
		short:       "ha",
		long:        "hectare",
		value:       f64(10000),
		dimension:   AREA,
		description: "SI-accepted unit (1 ha = 10,000 m²)",
	},
}
