package unit

import "math/big"

var pressureUnits = []baseUnit{
	{
		name:        "Pa",
		aliases:     []string{"pascal"},
		value:       f64(1),
		dimension:   PRESSURE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
	{
		name:        "psi",
		value:       f64(6894.757),
		dimension:   PRESSURE,
		description: "US/Imperial unit (pound per square inch)",
	},
	{
		name:      "at",
		aliases:   []string{"technical-atmosphere"},
		value:     f64(98066.5),
		dimension: PRESSURE,
	},
	{
		name:      "atm",
		aliases:   []string{"atmosphere", "standard-atmosphere"},
		value:     f64(101325),
		dimension: PRESSURE,
	},
	{
		name:        "bar",
		aliases:     []string{"bar"},
		value:       f64(100000),
		dimension:   PRESSURE,
		prefixes:    &metricPrefixes,
	},
	{
		name:      "torr",
		aliases:   []string{"Torr"},
		value:     big.NewRat(101325, 760),
		dimension: PRESSURE,
	},
	{
		name:      "mmHg",
		value:     f64(133.322387415),
		dimension: PRESSURE,
	},
	{
		name:      "mmH2O",
		value:     f64(9.80665),
		dimension: PRESSURE,
	},
	{
		name:      "cmH2O",
		value:     f64(98.0665),
		dimension: PRESSURE,
	},
}
