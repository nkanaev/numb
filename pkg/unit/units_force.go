package unit

var forceUnits = []baseUnit{
	{
		short:       "N",
		long:        "newton",
		value:       f64(1),
		dimension:   FORCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		short:     "dyn",
		long:      "dyne",
		value:     exp(10, -5),
		dimension: FORCE,
	},
	{
		short:     "lbf",
		long:      "poundforce",
		value:     f64(4.4482216152605),
		dimension: FORCE,
	},
	{
		short:     "kip",
		value:     f64(4448.2216),
		dimension: FORCE,
	},
	{
		short:     "pdl",
		long:      "poundal",
		value:     f64(0.138254954376),
		dimension: FORCE,
	},
}
