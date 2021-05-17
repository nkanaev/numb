package unit

var powerUnits = []baseUnit{
	{name: "W", long: "watt", value: f64(1), dimension: POWER, prefixes: &metricPrefixes, info: "SI derived unit"},
	{name: "hp", long: "horsepower", value: f64(745.69987158227), dimension: POWER, info: "mechanical horsepower"},
}
