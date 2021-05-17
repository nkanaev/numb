package unit

import "math/big"

var temperatureUnits = []baseUnit{
	{
		short:       "K",
		long:        "kelvin",
		value:       f64(1),
		dimension:   TEMPERATURE,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
	{
		short:       "°C, degC",
		long:        "celsius",
		value:       f64(1),
		offset:      f64(273.15),
		dimension:   TEMPERATURE,
		description: "SI derived unit",
	},
	{
		short:     "°F, degF",
		long:      "fahrenheit",
		value:     big.NewRat(10, 18),
		offset:    f64(459.67),
		dimension: TEMPERATURE,
	},
}
