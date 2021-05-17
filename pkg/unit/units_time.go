package unit

var timeUnits = []baseUnit{
	{
		short:       "s",
		long:        "sec, second",
		value:       f64(1),
		dimension:   TIME,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
	{
		short:       "min",
		long:        "minute",
		value:       f64(60),
		dimension:   TIME,
		description: "SI-accepted unit (1 min = 60 s)",
	},
	{
		short:       "h",
		long:        "hr, hour",
		value:       f64(3600),
		dimension:   TIME,
		description: "SI-accepted unit (1 h = 60 min)",
	},
	{
		short:       "d",
		long:        "day",
		value:       f64(86400),
		dimension:   TIME,
		description: "SI-accepted unit (1 day = 24 h)",
	},
	{
		short:     "week",
		value:     f64(7 * 86400),
		dimension: TIME,
	},
	{
		short:       "month",
		value:       f64(2629800),
		dimension:   TIME,
		description: "1/12th of Julian Year",
	},
	{
		short:       "year",
		value:       f64(31557600),
		dimension:   TIME,
		description: "Julian Year (365.25 days)",
	},
	{
		short:       "decade",
		value:       f64(315576000),
		dimension:   TIME,
		description: "Julian decade",
	},
	{
		short:       "century",
		value:       f64(3155760000),
		dimension:   TIME,
		description: "Julian century",
	},
	{
		short:       "millenium",
		value:       f64(31557600000),
		dimension:   TIME,
		description: "Julian millenium",
	},
}
