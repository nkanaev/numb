package unit

var forceUnits = []baseUnit{
	{
		name:      "N",
		aliases:   []string{"newton"},
		value:     f64(1),
		dimension: FORCE,
		prefixes:  &metricPrefixes,
		description: "SI derived unit",
	},
	{
		name:      "dyn",
		aliases:   []string{"dyne"},
		value:     exp(10, -5),
		dimension: FORCE,
	},
	{
		name:      "lbf",
		aliases:   []string{"poundforce"},
		value:     f64(4.4482216152605),
		dimension: FORCE,
	},
	{
		name:      "kip",
		value:     f64(4448.2216),
		dimension: FORCE,
	},
	{
		name:      "pdl",
		aliases:   []string{"poundal"},
		value:     f64(0.138254954376),
		dimension: FORCE,
	},
}
