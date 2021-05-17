package unit

var lengthUnits = []baseUnit{
	{d: LENGTH, name: "m", long: "meter, metre", value: f64(1), prefixes: &metricPrefixes, info: "SI base unit"},
	{d: LENGTH, name: "in", long: "inch", value: f64(0.0254)},
	{d: LENGTH, name: "ft", long: "foot, feet", value: f64(0.3048)},
	{d: LENGTH, name: "yd", long: "yard, yards", value: f64(0.9144)},
	{d: LENGTH, name: "mi", long: "mile, miles", value: f64(1609.344)},

	{d: LENGTH, name: "li", long: "link", value: f64(0.201168)},
	{d: LENGTH, name: "rd", long: "rod", value: f64(5.0292)},
	{d: LENGTH, name: "ch", long: "chain", value: f64(20.1)},
	{d: LENGTH, name: "angstrom", value: exp(10, -10)},
	{d: LENGTH, name: "mil", value: f64(0.0000254)},
	{d: LENGTH, name: "au", long: "astronomical-unit", value: f64(149597870700), info: "SI-accepted unit"},
	{d: LENGTH, name: "ly", long: "lightyeaar, light-year", value: f64(9460730472580800), prefixes: &metricPrefixes, info: "SI-accepted unit"},
	{d: LENGTH, name: "lightsecond", long: "lightsecond, light-second", value: f64(299792458), prefixes: &metricPrefixes},
}
