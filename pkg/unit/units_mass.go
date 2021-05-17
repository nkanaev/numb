package unit

var massUnits = []baseUnit{
	{d: MASS, name: "g", long: "gram", value: f64(0.001), prefixes: &metricPrefixes, info: "(0.001 kg - SI base unit)"},
	{d: MASS, name: "t", long: "tonne, metric-ton", value: f64(1000), prefixes: &metricPrefixes, info: "SI-accepted unit (1 t = 1000 kg)"},
	{d: MASS, name: "Da", long: "dalton", value: parse("1.6605402e-27"), info: "SI-accepted unit"},
}
