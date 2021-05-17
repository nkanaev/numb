package unit

var timeUnits = []baseUnit{
	{d: TIME, name: "s", long: "sec, second", value: f64(1), prefixes: &metricPrefixes, info: "SI base unit"},
	{d: TIME, name: "min", long: "minute", value: f64(60), info: "SI-accepted unit (1 min = 60 s)"},
	{d: TIME, name: "h", long: "hr, hour", value: f64(3600), info: "SI-accepted unit (1 h = 60 min)"},
	{d: TIME, name: "d", long: "day", value: f64(86400), info: "SI-accepted unit (1 day = 24 h)"},
	{d: TIME, name: "week", value: f64(7 * 86400)},
	{d: TIME, name: "month", value: f64(2629800), info: "1/12th of Julian Year"},
	{d: TIME, name: "year", value: f64(31557600), info: "Julian Year (365.25 days)"},
	{d: TIME, name: "decade", value: f64(315576000), info: "Julian decade"},
	{d: TIME, name: "century", value: f64(3155760000), info: "Julian century"},
	{d: TIME, name: "millenium", value: f64(31557600000), info: "Julian millenium"},
}
