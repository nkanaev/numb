package unit

var volumeUnits = []baseUnit{
	{d: VOLUME, name: "m³, m3", value: f64(1), prefixes: &metricPrefixes, prefixpow: 3},
	{d: VOLUME, name: "l, lt", long: "liter, litre", value: f64(0.001), prefixes: &metricPrefixes, info: "SI-accepted unit (1 l = 0.001 m³)"},
	{d: VOLUME, name: "in³, in3, cuin", value: f64(1.6387064e-5)},
	{d: VOLUME, name: "ft³, ft3, cuft", value: f64(0.028316846592)},
	{d: VOLUME, name: "yd³, yd3, cuyd", value: f64(0.764554857984)},
	{d: VOLUME, name: "teaspoon", value: f64(0.000005)},
	{d: VOLUME, name: "tablespoon", value: f64(0.000015)},
}
