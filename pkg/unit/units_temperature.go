package unit

import "math/big"

var temperatureUnits = []baseUnit{
	{
		name:      "K",
		aliases:   []string{"kelvin"},
		value:     f64(1),
		dimension: TEMPERATURE,
		prefixes:  &metricPrefixes,
	},
	{
		name:         "°C",
		shortaliases: []string{"degC"},
		aliases:      []string{"celsius"},
		value:        f64(1),
		offset:       f64(273.15),
		dimension:    TEMPERATURE,
	},
	{
		name:         "°F",
		shortaliases: []string{"degF"},
		aliases:      []string{"fahrenheit"},
		value:        big.NewRat(10, 18),
		offset:       f64(459.67),
		dimension:    TEMPERATURE,
	},
}
