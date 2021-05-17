package unit

var forceUnits = []baseUnit{
	{d: FORCE, name: "N", long: "newton", value: f64(1), prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: FORCE, name: "dyn", long: "dyne", value: exp(10, -5)},
	{d: FORCE, name: "lbf", long: "poundforce", value: f64(4.4482216152605)},
	{d: FORCE, name: "kip", value: f64(4448.2216)},
	{d: FORCE, name: "pdl", long: "poundal", value: f64(0.138254954376)},
}
