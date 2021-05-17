package unit

var energyUnits = []baseUnit{
	{d: ENERGY, name: "J", long: "joule", value: f64(1), prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ENERGY, name: "Wh", long: "watt-hour", value: f64(3600), prefixes: &metricPrefixes},
	{d: ENERGY, name: "eV", long: "electronvolt", value: f64(1.602176565e-19), prefixes: &metricPrefixes, info: "SI-accepted unit"},
	{d: ENERGY, name: "erg", value: exp(10, -7)},
}
