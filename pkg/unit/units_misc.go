package unit

var frequencyUnits = []baseUnit{
	{name: "Hz", long: "hertz", value: one, dimension: FREQUENCY, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricCurrentUnits = []baseUnit{
	{name: "A, amp", long: "ampere", value: one, dimension: ELECTRIC_CURRENT, prefixes: &metricPrefixes, info: "SI base unit"},
}

var luminousIntensityUnits = []baseUnit{
	{name: "cd", long: "candela", value: one, dimension: LUMINOUS_INTENSITY, prefixes: &metricPrefixes, info: "SI base unit"},
}

var amountOfSubstanceUnits = []baseUnit{
	{name: "mol", long: "mole", value: one, dimension: AMOUNT_OF_SUBSTANCE, prefixes: &metricPrefixes, info: "SI base unit"},
}

var electricChargeUnits = []baseUnit{
	{name: "C", long: "coulomb", value: one, dimension: ELECTRIC_CHARGE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricPotentialUnits = []baseUnit{
	{name: "V", long: "volt", value: one, dimension: ELECTRIC_POTENTIAL, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricCapaticanceUnits = []baseUnit{
	{name: "F", long: "farad", value: one, dimension: ELECTRIC_CAPACITANCE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricConductanceUnits = []baseUnit{
	{name: "S", long: "siemens", value: one, dimension: ELECTRIC_CONDUCTANCE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var magneticFluxUnits = []baseUnit{
	{name: "Wb", long: "weber", value: one, dimension: MAGNETIC_FLUX, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var magneticFluxDensityUnits = []baseUnit{
	{name: "T", long: "tesla", value: one, dimension: MAGNETIC_FLUX_DENSITY, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricInductanceUnits = []baseUnit{
	{name: "H", long: "henry", value: one, dimension: ELECTRIC_INDUCTANCE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var electricResistanceUnits = []baseUnit{
	{name: "Ω, ohm", long: "ohm", value: one, dimension: ELECTRIC_RESISTANCE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var solidAngleUnits = []baseUnit{
	{name: "sr", long: "steradian", value: one, dimension: SOLID_ANGLE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var ionizingRadiationUnits = []baseUnit{
	{name: "Sv", long: "sievert", value: one, dimension: IONIZING_RADIATION, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var radiationDoseUnits = []baseUnit{
	{name: "Gy", long: "gray", value: one, dimension: RADIATION_DOSE, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var catalycticActivityUnits = []baseUnit{
	{name: "kat", long: "katal", value: one, dimension: CATALYCTIC_ACTIVITY, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var luminousFluxUnits = []baseUnit{
	{name: "lm", long: "lumen", value: one, dimension: LUMINOUS_FLUX, prefixes: &metricPrefixes, info: "SI derived unit"},
}

var illuminanceUnits = []baseUnit{
	{name: "lx", long: "lux", value: one, dimension: ILLUMINANCE, prefixes: &metricPrefixes, info: "SI derived unit"},
}
