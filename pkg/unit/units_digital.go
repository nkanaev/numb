package unit

var digitalUnits = []baseUnit{
	{
		short:     "bit",
		long:      "bit",
		value:     f64(1),
		dimension: DIGITAL,
		prefixes:  &digitalPrefixes,
	},
	{
		short:     "b",
		long:      "byte",
		value:     f64(8),
		dimension: DIGITAL,
		prefixes:  &digitalPrefixes,
	},
}
