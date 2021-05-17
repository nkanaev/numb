package unit

var powerUnits = []baseUnit{
	{d: POWER, name: "W", long: "watt", value: f64(1), prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: POWER, name: "hp", long: "horsepower", value: f64(745.69987158227), info: "mechanical horsepower"},
}
