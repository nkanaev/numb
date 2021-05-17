package unit

var forceUnits = []baseUnit{
	{name: "N", long: "newton", value: f64(1), dimension: FORCE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{name: "dyn", long: "dyne", value: exp(10, -5), dimension: FORCE},
	{name: "lbf", long: "poundforce", value: f64(4.4482216152605), dimension: FORCE},
	{name: "kip", value: f64(4448.2216), dimension: FORCE},
	{name: "pdl", long: "poundal", value: f64(0.138254954376), dimension: FORCE},
}
