package unit

import "math/big"

var pressureUnits = []baseUnit{
	{d: PRESSURE, name: "Pa", long: "pascal", value: f64(1), prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: PRESSURE, name: "psi", value: f64(6894.757), info: "US/Imperial unit (pound per square inch)"},
	{d: PRESSURE, name: "at", long: "technical-atmosphere", value: f64(98066.5)},
	{d: PRESSURE, name: "atm", long: "atmosphere, standard-atmosphere", value: f64(101325)},
	{d: PRESSURE, name: "bar", long: "bar", value: f64(100000), prefixes: &metricPrefixes},
	{d: PRESSURE, name: "torr", long: "Torr", value: big.NewRat(101325, 760)},
	{d: PRESSURE, name: "mmHg", value: f64(133.322387415)},
	{d: PRESSURE, name: "mmH2O", value: f64(9.80665)},
	{d: PRESSURE, name: "cmH2O", value: f64(98.0665)},
}
