package unit

var digitalUnits = []baseUnit{
	{name: "bit", long: "bit", value: f64(1), dimension: DIGITAL, prefixes: &digitalPrefixes},
	{name: "b", long: "byte", value: f64(8), dimension: DIGITAL, prefixes: &digitalPrefixes},
}
