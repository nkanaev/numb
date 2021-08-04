package unit

import (
	"math/big"

	r "github.com/nkanaev/numb/pkg/ratutils"
	d "github.com/nkanaev/numb/pkg/unit/dimension"
)

// International Yard & Pound
// US 1959 / AU 1964 / UK 1964
// https://en.wikipedia.org/wiki/International_yard_and_pound
var iyard = r.Num("0.9144")
var ipound = r.Num("0.45359237")

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

	{u: d.ENERGY, name: "J", long: "joule", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ENERGY, name: "Wh", long: "watt-hour", value: r.Num("3600"), prefixes: &metricPrefixes},
	{u: d.ENERGY, name: "eV", long: "electronvolt", value: r.Num("1.602176565e-19"), prefixes: &metricPrefixes, info: "accepted for use with the SI"},
	{u: d.ENERGY, name: "erg", value: r.Exp(10, -7)},

	{u: d.FORCE, name: "N", long: "newton", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.FORCE, name: "dyn", long: "dyne", value: r.Exp(10, -5)},
	{u: d.FORCE, name: "lbf", long: "poundforce", value: r.Num("4.4482216152605")},
	{u: d.FORCE, name: "kip", value: r.Num("4448.2216")},
	{u: d.FORCE, name: "pdl", long: "poundal", value: r.Num("0.138254954376")},

	{u: d.PRESSURE, name: "Pa", long: "pascal", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.PRESSURE, name: "psi", value: r.Num("6894.757"), info: "US/Imperial unit (pound per square inch)"},
	{u: d.PRESSURE, name: "at", long: "technical-atmosphere", value: r.Num("98066.5")},
	{u: d.PRESSURE, name: "atm", long: "atmosphere, standard-atmosphere", value: r.Num("101325")},
	{u: d.PRESSURE, name: "bar", long: "bar", value: r.Num("100000"), prefixes: &metricPrefixes},
	{u: d.PRESSURE, name: "torr", long: "Torr", value: r.Num("101325/760")},
	{u: d.PRESSURE, name: "mmHg", value: r.Num("133.322387415")},
	{u: d.PRESSURE, name: "mmH2O", value: r.Num("9.80665")},
	{u: d.PRESSURE, name: "cmH2O", value: r.Num("98.0665")},

	{u: d.POWER, name: "W", long: "watt", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.POWER, name: "hp", long: "horsepower", value: r.Num("745.69987158227"), info: "mechanical horsepower"},

	{u: d.RADIOACTIVITY, name: "Bq", long: "becquierel", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.RADIOACTIVITY, name: "Ci", long: "curie", value: r.Num("3.7e10")},
	{u: d.RADIOACTIVITY, name: "Rd", long: "rutherford", value: r.Exp(10, 6)},

	{u: d.AMOUNT_OF_SUBSTANCE, name: "mol", long: "mole", value: r.ONE, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: d.CATALYCTIC_ACTIVITY, name: "kat", long: "katal", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_CAPACITANCE, name: "F", long: "farad", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_CHARGE, name: "C", long: "coulomb", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_CONDUCTANCE, name: "S", long: "siemens", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_CURRENT, name: "A, amp", long: "ampere", value: r.ONE, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: d.ELECTRIC_INDUCTANCE, name: "H", long: "henry", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_POTENTIAL, name: "V", long: "volt", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ELECTRIC_RESISTANCE, name: "ohm", long: "ohm", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.FREQUENCY, name: "Hz", long: "hertz", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.ILLUMINANCE, name: "lx", long: "lux", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.IONIZING_RADIATION, name: "Sv", long: "sievert", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.LUMINOUS_FLUX, name: "lm", long: "lumen", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.LUMINOUS_INTENSITY, name: "cd", long: "candela", value: r.ONE, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: d.MAGNETIC_FLUX, name: "Wb", long: "weber", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.MAGNETIC_FLUX_DENSITY, name: "T", long: "tesla", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.RADIATION_DOSE, name: "Gy", long: "gray", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: d.SOLID_ANGLE, name: "sr", long: "steradian", value: r.ONE, prefixes: &metricPrefixes, info: "SI derived unit"},

	// speed: base is meter/sec
	{u: d.SPEED, name: "kmph", value: big.NewRat(5, 18)},
	{u: d.SPEED, name: "mph", value: big.NewRat(1397, 3125)},
	{u: d.SPEED, name: "kn, knot", value: big.NewRat(463, 900)},

	// data-rate: base is bit/sec
	{u: d.DATA_RATE, name: "kbps", value: r.Exp(10, 3)},
	{u: d.DATA_RATE, name: "Mbps", value: r.Exp(10, 6)},
	{u: d.DATA_RATE, name: "Gbps", value: r.Exp(10, 9)},
	{u: d.DATA_RATE, name: "Tbps", value: r.Exp(10, 12)},

	{u: d.CURRENCY, name: "CURRENCY", value: r.ONE},
}
