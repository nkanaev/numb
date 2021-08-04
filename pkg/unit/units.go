package unit

import (
	r "github.com/nkanaev/numb/pkg/ratutils"
	d "github.com/nkanaev/numb/pkg/unit/dimension"
)

var units = []unitDef{
	{u: d.LENGTH, name: "m", long: "meter, metre", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.MASS, name: "g", long: "gram", value: r.Num("0.001"), prefixes: &metricPrefixes},
	{u: d.TIME, name: "s", long: "sec, second", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.TEMPERATURE, name: "K", long: "kelvin", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.TEMPERATURE, name: "°C, degC", long: "celsius", value: r.ONE, offset: r.Num("273.15")},
	{u: d.TEMPERATURE, name: "°F, degF", long: "fahrenheit", value: r.Num("10/18"), offset: r.Num("459.67")},
	{u: d.ANGLE, name: "rad", long: "radian", value: r.ONE},
	{u: d.DIGITAL, name: "bit", long: "bit", value: r.ONE, prefixes: &digitalPrefixes},
	{u: d.DIGITAL, name: "B", long: "byte", value: r.Num("8"), prefixes: &digitalPrefixes},
	{u: d.AREA, name: "m2", value: r.ONE, prefixes: &metricPrefixes, prefixpow: 2},
	{u: d.VOLUME, name: "m3", value: r.ONE, prefixes: &metricPrefixes, prefixpow: 3},
	{u: d.VOLUME, name: "l, lt", long: "liter, litre", value: r.Num("0.001"), prefixes: &metricPrefixes},
	{u: d.ENERGY, name: "J", long: "joule", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ENERGY, name: "Wh", long: "watt-hour", value: r.Num("3600"), prefixes: &metricPrefixes},
	{u: d.ENERGY, name: "eV", long: "electronvolt", value: r.Num("1.602176565e-19"), prefixes: &metricPrefixes},
	{u: d.FORCE, name: "N", long: "newton", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.PRESSURE, name: "Pa", long: "pascal", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.PRESSURE, name: "bar", long: "bar", value: r.Num("100000"), prefixes: &metricPrefixes},
	{u: d.POWER, name: "W", long: "watt", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.RADIOACTIVITY, name: "Bq", long: "becquierel", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.AMOUNT_OF_SUBSTANCE, name: "mol", long: "mole", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.CATALYCTIC_ACTIVITY, name: "kat", long: "katal", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_CAPACITANCE, name: "F", long: "farad", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_CHARGE, name: "C", long: "coulomb", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_CONDUCTANCE, name: "S", long: "siemens", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_CURRENT, name: "A, amp", long: "ampere", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_INDUCTANCE, name: "H", long: "henry", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_POTENTIAL, name: "V", long: "volt", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ELECTRIC_RESISTANCE, name: "ohm", long: "ohm", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.FREQUENCY, name: "Hz", long: "hertz", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.ILLUMINANCE, name: "lx", long: "lux", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.IONIZING_RADIATION, name: "Sv", long: "sievert", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.LUMINOUS_FLUX, name: "lm", long: "lumen", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.LUMINOUS_INTENSITY, name: "cd", long: "candela", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.MAGNETIC_FLUX, name: "Wb", long: "weber", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.MAGNETIC_FLUX_DENSITY, name: "T", long: "tesla", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.RADIATION_DOSE, name: "Gy", long: "gray", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.SOLID_ANGLE, name: "sr", long: "steradian", value: r.ONE, prefixes: &metricPrefixes},
	{u: d.CURRENCY, name: "CURRENCY", value: r.ONE},
}
