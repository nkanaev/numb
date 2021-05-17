package unit

var radioactivityUnits = []baseUnit{
	{d: RADIOACTIVITY, name: "Bq", long: "becquierel", value: f64(1), prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: RADIOACTIVITY, name: "Ci", long: "curie", value: f64(3.7e+10)},
	{d: RADIOACTIVITY, name: "Rd", long: "rutherford", value: exp(10, 6)},
}
