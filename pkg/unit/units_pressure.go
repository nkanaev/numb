package unit

import "math/big"

var pressureUnits = []baseUnit{
	{
		short:       "Pa",
		long:        "pascal",
		value:       f64(1),
		dimension:   PRESSURE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		short:       "psi",
		value:       f64(6894.757),
		dimension:   PRESSURE,
		description: "US/Imperial unit (pound per square inch)",
	},
	{
		short:     "at",
		long:      "technical-atmosphere",
		value:     f64(98066.5),
		dimension: PRESSURE,
	},
	{
		short:     "atm",
		long:      "atmosphere, standard-atmosphere",
		value:     f64(101325),
		dimension: PRESSURE,
	},
	{
		short:     "bar",
		long:      "bar",
		value:     f64(100000),
		dimension: PRESSURE,
		prefixes:  &metricPrefixes,
	},
	{
		short:     "torr",
		long:      "Torr",
		value:     big.NewRat(101325, 760),
		dimension: PRESSURE,
	},
	{
		short:     "mmHg",
		value:     f64(133.322387415),
		dimension: PRESSURE,
	},
	{
		short:     "mmH2O",
		value:     f64(9.80665),
		dimension: PRESSURE,
	},
	{
		short:     "cmH2O",
		value:     f64(98.0665),
		dimension: PRESSURE,
	},
}
