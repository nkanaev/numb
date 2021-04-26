package unit

var digitalUnits = []baseUnit{
	{
		name:      "bit",
		aliases:   []string{"bit"},
		value:     f64(1),
		dimension: DIGITAL,
		prefixes:  &digitalPrefixes,
	},
	{
		name:      "b",
		aliases:   []string{"byte"},
		value:     f64(8),
		dimension: DIGITAL,
		prefixes:  &digitalPrefixes,
	},
}
