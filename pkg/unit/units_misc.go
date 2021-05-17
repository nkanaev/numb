package unit

var frequencyUnits = []baseUnit{
	{d: FREQUENCY, name: "Hz", long: "hertz", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricCurrentUnits = []baseUnit{
	{d: ELECTRIC_CURRENT, name: "A, amp", long: "ampere", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
}

var luminousIntensityUnits = []baseUnit{
	{d: LUMINOUS_INTENSITY, name: "cd", long: "candela", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
}

var amountOfSubstanceUnits = []baseUnit{
	{d: AMOUNT_OF_SUBSTANCE, name: "mol", long: "mole", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
}

var electricChargeUnits = []baseUnit{
	{d: ELECTRIC_CHARGE, name: "C", long: "coulomb", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricPotentialUnits = []baseUnit{
	{d: ELECTRIC_POTENTIAL, name: "V", long: "volt", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricCapaticanceUnits = []baseUnit{
	{d: ELECTRIC_CAPACITANCE, name: "F", long: "farad", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricConductanceUnits = []baseUnit{
	{d: ELECTRIC_CONDUCTANCE, name: "S", long: "siemens", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var magneticFluxUnits = []baseUnit{
	{d: MAGNETIC_FLUX, name: "Wb", long: "weber", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var magneticFluxDensityUnits = []baseUnit{
	{d: MAGNETIC_FLUX_DENSITY, name: "T", long: "tesla", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricInductanceUnits = []baseUnit{
	{d: ELECTRIC_INDUCTANCE, name: "H", long: "henry", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricResistanceUnits = []baseUnit{
	{d: ELECTRIC_RESISTANCE, name: "Ω, ohm", long: "ohm", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var solidAngleUnits = []baseUnit{
	{d: SOLID_ANGLE, name: "sr", long: "steradian", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var ionizingRadiationUnits = []baseUnit{
	{d: IONIZING_RADIATION, name: "Sv", long: "sievert", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var radiationDoseUnits = []baseUnit{
	{d: RADIATION_DOSE, name: "Gy", long: "gray", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var catalycticActivityUnits = []baseUnit{
	{d: CATALYCTIC_ACTIVITY, name: "kat", long: "katal", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var luminousFluxUnits = []baseUnit{
	{d: LUMINOUS_FLUX, name: "lm", long: "lumen", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var illuminanceUnits = []baseUnit{
	{d: ILLUMINANCE, name: "lx", long: "lux", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}
