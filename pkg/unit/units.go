package unit

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/consts"
)

var one = big.NewRat(1, 1)

func div(a *big.Rat, x int64) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, big.NewRat(x, 1))
	return num
}

func divr(a, x *big.Rat) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, x)
	return num
}

func num(x string) *big.Rat {
	rat, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse: " + x)
	}
	return rat
}

func mul(a *big.Rat, n int64) *big.Rat {
	x := new(big.Rat).Set(a)
	x.Mul(x, big.NewRat(n, 1))
	return x
}

// International Yard & Pound
// US 1959 / AU 1964 / UK 1964
// https://en.wikipedia.org/wiki/International_yard_and_pound
var iyard = num("0.9144")
var ipound = num("0.45359237")

var units = []unitDef{
	{u: LENGTH, name: "m", long: "meter, metre", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	// lengths: US & Imperial
	{u: LENGTH, name: "in, inch", value: div(iyard, 3*12), info: "1/12 feet"},
	{u: LENGTH, name: "ft, foot, feet", value: div(iyard, 3), info: "1/3 yard"},
	{u: LENGTH, name: "yd, yard", value: iyard, info: "International Yard (3 feet)"},
	{u: LENGTH, name: "mi, mile", value: mul(iyard, 22*10*8), info: "8 furlongs"},
	// lengths: misc
	{u: LENGTH, name: "angstrom", value: exp(10, -10)},
	{u: LENGTH, name: "au", long: "astronomical-unit", value: num("149597870700"), info: "accepted for use with the SI"},
	{u: LENGTH, name: "pc, parsec", value: divr(mul(num("149597870700"), 648000), consts.PI), info: "648000/pi astronomical units"},
	{u: LENGTH, name: "ly", long: "lightyear, light-year", value: num("9460730472580800"), prefixes: &metricPrefixes, info: "accepted for use with the SI"},
	{u: LENGTH, name: "lightsecond", long: "lightsecond, light-second", value: num("299792458"), prefixes: &metricPrefixes},

	{u: MASS, name: "g", long: "gram", value: num("0.001"), prefixes: &metricPrefixes, info: "(0.001 kg - SI base unit)"},
	{u: MASS, name: "t", long: "tonne, metric-ton", value: num("1000"), prefixes: &metricPrefixes, info: "accepted for use with the SI (1 t = 1000 kg)"},
	{u: MASS, name: "Da", long: "dalton", value: num("1.6605402e-27"), info: "accepted for use with the SI"},
	// avoirdupois system
	{u: MASS, name: "dr, dram", value: div(ipound, 256), info: "1/256 pound"},
	{u: MASS, name: "oz, once, ounce", value: div(ipound, 16), info: "1/16 pound"},
	{u: MASS, name: "lb, pound", value: ipound, info: "International pound"},

	{u: TIME, name: "s", long: "sec, second", value: num("1"), prefixes: &metricPrefixes, info: "SI base unit"},
	{u: TIME, name: "min", long: "minute", value: num("60"), info: "accepted for use with the SI (1 min = 60 s)"},
	{u: TIME, name: "h", long: "hr, hour", value: num("3600"), info: "accepted for use with the SI (1 h = 60 min)"},
	{u: TIME, name: "d", long: "day", value: num("86400"), info: "accepted for use with the SI (1 day = 24 h)"},
	{u: TIME, name: "week", value: num("604800")},
	{u: TIME, name: "month", value: num("2629800"), info: "1/12th of Julian Year"},
	{u: TIME, name: "year", value: num("31557600"), info: "Julian Year (365.25 days)"},
	{u: TIME, name: "decade", value: num("315576000"), info: "Julian decade"},
	{u: TIME, name: "century", value: num("3155760000"), info: "Julian century"},
	{u: TIME, name: "millenium", value: num("31557600000"), info: "Julian millenium"},

	{u: TEMPERATURE, name: "K", long: "kelvin", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: TEMPERATURE, name: "°C, degC", long: "celsius", value: one, offset: num("273.15"), info: "SI derived unit"},
	{u: TEMPERATURE, name: "°F, degF", long: "fahrenheit", value: num("10/18"), offset: num("459.67")},

	{u: ANGLE, name: "rad", long: "radian", value: one, info: "SI derived unit"},
	{u: ANGLE, name: "°", long: "deg, degree", value: div(consts.PI, 180), info: "accepted for use with the SI"},
	{u: ANGLE, name: "arcsec", value: div(consts.PI, 648000), info: "accepted for use with the SI (pi / 648000)"},
	{u: ANGLE, name: "arcmin", value: div(consts.PI, 10800), info: "accepted for use with the SI (pi / 10800)"},
	{u: ANGLE, name: "grad", long: "grade, gradian", value: div(consts.PI, 200)},
	{u: ANGLE, name: "cycle", value: div(consts.PI, 2)},

	{u: DIGITAL, name: "bit", long: "bit", value: one, prefixes: &digitalPrefixes},
	{u: DIGITAL, name: "B", long: "byte", value: num("8"), prefixes: &digitalPrefixes},

	{u: AREA, name: "m², m2", value: one, prefixes: &metricPrefixes, prefixpow: 2},
	{u: AREA, name: "in², in2, sqin", value: num("0.00064516")},
	{u: AREA, name: "ft², ft2, sqft", long: "sqfeet", value: num("0.09290304")},
	{u: AREA, name: "yd², sqyd", long: "sqyard", value: num("0.83612736")},
	{u: AREA, name: "rd², rd2, sqrd", value: num("25.29295")},
	{u: AREA, name: "ch², sqch", value: num("404.6873")},
	{u: AREA, name: "mi², sqmil, sqmi", long: "sqmile", value: num("6.4516e-10")},
	{u: AREA, name: "acre", value: num("4046.86")},
	{u: AREA, name: "ha", long: "hectare", value: num("10000"), info: "accepted for use with the SI (1 ha = 10,000 m²)"},

	{u: VOLUME, name: "m³, m3", value: one, prefixes: &metricPrefixes, prefixpow: 3},
	{u: VOLUME, name: "l, lt", long: "liter, litre", value: num("0.001"), prefixes: &metricPrefixes, info: "accepted for use with the SI (1 l = 0.001 m³)"},
	{u: VOLUME, name: "in³, in3, cuin", value: num("1.6387064e-5")},
	{u: VOLUME, name: "ft³, ft3, cuft", value: num("0.028316846592")},
	{u: VOLUME, name: "yd³, yd3, cuyd", value: num("0.764554857984")},

	{u: VOLUME, name: "teaspoon", value: num("0.000005")},
	{u: VOLUME, name: "tablespoon", value: num("0.000015")},

	{u: ENERGY, name: "J", long: "joule", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ENERGY, name: "Wh", long: "watt-hour", value: num("3600"), prefixes: &metricPrefixes},
	{u: ENERGY, name: "eV", long: "electronvolt", value: num("1.602176565e-19"), prefixes: &metricPrefixes, info: "accepted for use with the SI"},
	{u: ENERGY, name: "erg", value: exp(10, -7)},

	{u: FORCE, name: "N", long: "newton", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: FORCE, name: "dyn", long: "dyne", value: exp(10, -5)},
	{u: FORCE, name: "lbf", long: "poundforce", value: num("4.4482216152605")},
	{u: FORCE, name: "kip", value: num("4448.2216")},
	{u: FORCE, name: "pdl", long: "poundal", value: num("0.138254954376")},

	{u: PRESSURE, name: "Pa", long: "pascal", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: PRESSURE, name: "psi", value: num("6894.757"), info: "US/Imperial unit (pound per square inch)"},
	{u: PRESSURE, name: "at", long: "technical-atmosphere", value: num("98066.5")},
	{u: PRESSURE, name: "atm", long: "atmosphere, standard-atmosphere", value: num("101325")},
	{u: PRESSURE, name: "bar", long: "bar", value: num("100000"), prefixes: &metricPrefixes},
	{u: PRESSURE, name: "torr", long: "Torr", value: num("101325/760")},
	{u: PRESSURE, name: "mmHg", value: num("133.322387415")},
	{u: PRESSURE, name: "mmH2O", value: num("9.80665")},
	{u: PRESSURE, name: "cmH2O", value: num("98.0665")},

	{u: POWER, name: "W", long: "watt", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: POWER, name: "hp", long: "horsepower", value: num("745.69987158227"), info: "mechanical horsepower"},

	{u: RADIOACTIVITY, name: "Bq", long: "becquierel", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: RADIOACTIVITY, name: "Ci", long: "curie", value: num("3.7e10")},
	{u: RADIOACTIVITY, name: "Rd", long: "rutherford", value: exp(10, 6)},

	{u: AMOUNT_OF_SUBSTANCE, name: "mol", long: "mole", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: CATALYCTIC_ACTIVITY, name: "kat", long: "katal", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_CAPACITANCE, name: "F", long: "farad", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_CHARGE, name: "C", long: "coulomb", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_CONDUCTANCE, name: "S", long: "siemens", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_CURRENT, name: "A, amp", long: "ampere", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: ELECTRIC_INDUCTANCE, name: "H", long: "henry", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_POTENTIAL, name: "V", long: "volt", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ELECTRIC_RESISTANCE, name: "ohm", long: "ohm", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: FREQUENCY, name: "Hz", long: "hertz", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: ILLUMINANCE, name: "lx", long: "lux", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: IONIZING_RADIATION, name: "Sv", long: "sievert", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: LUMINOUS_FLUX, name: "lm", long: "lumen", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: LUMINOUS_INTENSITY, name: "cd", long: "candela", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{u: MAGNETIC_FLUX, name: "Wb", long: "weber", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: MAGNETIC_FLUX_DENSITY, name: "T", long: "tesla", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: RADIATION_DOSE, name: "Gy", long: "gray", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{u: SOLID_ANGLE, name: "sr", long: "steradian", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}
