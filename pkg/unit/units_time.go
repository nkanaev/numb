package unit

var timeUnits = []baseUnit{
	{
		name:      "s",
		aliases:   []string{"sec", "second"},
		value:     f64(1),
		dimension: TIME,
		prefixes:  &metricPrefixes,
	},
	{
		name:      "min",
		aliases:   []string{"minute"},
		value:     f64(60),
		dimension: TIME,
	},
	{
		name:      "h",
		aliases:   []string{"hr", "hour"},
		value:     f64(3600),
		dimension: TIME,
	},
	{
		name:      "day",
		aliases:   []string{},
		value:     f64(86400),
		dimension: TIME,
	},
	{
		name:      "week",
		aliases:   []string{},
		value:     f64(7 * 86400),
		dimension: TIME,
	},
	{
		name:      "month",
		aliases:   []string{},
		value:     f64(2629800),
		dimension: TIME,
	},
	{
		name:      "year",
		aliases:   []string{},
		value:     f64(31557600),
		dimension: TIME,
	},
}
