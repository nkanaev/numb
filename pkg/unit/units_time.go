package unit

var timeUnits = []baseUnit{
	{name: "s", long: "sec, second", value: f64(1), dimension: TIME, prefixes: &metricPrefixes, info: "SI base unit"},
	{name: "min", long: "minute", value: f64(60), dimension: TIME, info: "SI-accepted unit (1 min = 60 s)"},
	{name: "h", long: "hr, hour", value: f64(3600), dimension: TIME, info: "SI-accepted unit (1 h = 60 min)"},
	{name: "d", long: "day", value: f64(86400), dimension: TIME, info: "SI-accepted unit (1 day = 24 h)"},
	{name: "week", value: f64(7 * 86400), dimension: TIME},
	{name: "month", value: f64(2629800), dimension: TIME, info: "1/12th of Julian Year"},
	{name: "year", value: f64(31557600), dimension: TIME, info: "Julian Year (365.25 days)"},
	{name: "decade", value: f64(315576000), dimension: TIME, info: "Julian decade"},
	{name: "century", value: f64(3155760000), dimension: TIME, info: "Julian century"},
	{name: "millenium", value: f64(31557600000), dimension: TIME, info: "Julian millenium"},
}
