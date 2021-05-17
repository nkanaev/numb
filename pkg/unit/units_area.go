package unit

var areaUnits = []baseUnit{
	{d: AREA, name: "m², m2", value: f64(1), prefixes: &metricPrefixes, prefixpow: 2},
	{d: AREA, name: "in², in2, sqin", value: f64(0.00064516)},
	{d: AREA, name: "ft², ft2, sqft", long: "sqfeet", value: f64(0.09290304)},
	{d: AREA, name: "yd², sqyd", long: "sqyard", value: f64(0.83612736)},
	{d: AREA, name: "rd², rd2, sqrd", value: f64(25.29295)},
	{d: AREA, name: "ch², sqch", value: f64(404.6873)},
	{d: AREA, name: "mi², sqmil, sqmi", long: "sqmile", value: f64(6.4516e-10)},
	{d: AREA, name: "acre", value: f64(4046.86)},
	{d: AREA, name: "ha", long: "hectare", value: f64(10000), info: "SI-accepted unit (1 ha = 10,000 m²)"},
}
