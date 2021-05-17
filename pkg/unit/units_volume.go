package unit

var volumeUnits = []baseUnit{
	{name: "m³, m3", value: f64(1), dimension: VOLUME, prefixes: &metricPrefixes, prefixpow: 3},
	{name: "l, lt", long: "liter, litre", value: f64(0.001), dimension: VOLUME, prefixes: &metricPrefixes, info: "SI-accepted unit (1 l = 0.001 m³)"},
	{name: "in³, in3, cuin", value: f64(1.6387064e-5), dimension: VOLUME},
	{name: "ft³, ft3, cuft", value: f64(0.028316846592), dimension: VOLUME},
	{name: "yd³, yd3, cuyd", value: f64(0.764554857984), dimension: VOLUME},
	{name: "teaspoon", value: f64(0.000005), dimension: VOLUME},
	{name: "tablespoon", value: f64(0.000015), dimension: VOLUME},
}
