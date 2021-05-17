package unit

var digitalUnits = []baseUnit{
	{d: DIGITAL, name: "bit", long: "bit", value: f64(1), prefixes: &digitalPrefixes},
	{d: DIGITAL, name: "b", long: "byte", value: f64(8), prefixes: &digitalPrefixes},
}
