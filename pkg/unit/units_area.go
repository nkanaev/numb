package unit

var areaUnits = []baseUnit{
	{name: "m², m2", value: f64(1), dimension: AREA, prefixes: &metricPrefixes, prefixpow: 2},
	{name: "in², in2, sqin", value: f64(0.00064516), dimension: AREA},
	{name: "ft², ft2, sqft", long: "sqfeet", value: f64(0.09290304), dimension: AREA},
	{name: "yd², sqyd", long: "sqyard", value: f64(0.83612736), dimension: AREA},
	{name: "rd², rd2, sqrd", value: f64(25.29295), dimension: AREA},
	{name: "ch², sqch", value: f64(404.6873), dimension: AREA},
	{name: "mi², sqmil, sqmi", long: "sqmile", value: f64(6.4516e-10), dimension: AREA},
	{name: "acre", value: f64(4046.86), dimension: AREA},
	{name: "ha", long: "hectare", value: f64(10000), dimension: AREA, info: "SI-accepted unit (1 ha = 10,000 m²)"},
}
