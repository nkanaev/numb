package unit

import "math/big"

var temperatureUnits = []baseUnit{
	{
		name:        "K",
		long:        "kelvin",
		value:       f64(1),
		dimension:   TEMPERATURE,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
	{
		name:        "°C, degC",
		long:        "celsius",
		value:       f64(1),
		offset:      f64(273.15),
		dimension:   TEMPERATURE,
		description: "SI derived unit",
	},
	{
		name:      "°F, degF",
		long:      "fahrenheit",
		value:     big.NewRat(10, 18),
		offset:    f64(459.67),
		dimension: TEMPERATURE,
	},
}
