package unit

var radioactivityUnits = []baseUnit{
	{name: "Bq", long: "becquierel", value: f64(1), dimension: RADIOACTIVITY, prefixes: &metricPrefixes, info: "SI derived unit"},
	{name: "Ci", long: "curie", value: f64(3.7e+10), dimension: RADIOACTIVITY},
	{name: "Rd", long: "rutherford", value: exp(10, 6), dimension: RADIOACTIVITY},
}
