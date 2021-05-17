package unit

import "math/big"

var temperatureUnits = []baseUnit{
	{d: TEMPERATURE, name: "K", long: "kelvin", value: f64(1), prefixes: &metricPrefixes, info: "SI base unit"},
	{d: TEMPERATURE, name: "°C, degC", long: "celsius", value: f64(1), offset: f64(273.15), info: "SI derived unit"},
	{d: TEMPERATURE, name: "°F, degF", long: "fahrenheit", value: big.NewRat(10, 18), offset: f64(459.67)},
}
