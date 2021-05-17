package unit

var frequencyUnits = []baseUnit{
	{
		name:        "Hz",
		long:        "hertz",
		value:       f64(1),
		dimension:   FREQUENCY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCurrentUnits = []baseUnit{
	{
		name:        "A, amp",
		long:        "ampere",
		value:       f64(1),
		dimension:   ELECTRIC_CURRENT,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var luminousIntensityUnits = []baseUnit{
	{
		name:        "cd",
		long:        "candela",
		value:       f64(1),
		dimension:   LUMINOUS_INTENSITY,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var amountOfSubstanceUnits = []baseUnit{
	{
		name:        "mol",
		long:        "mole",
		value:       f64(1),
		dimension:   AMOUNT_OF_SUBSTANCE,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var electricChargeUnits = []baseUnit{
	{
		name:        "C",
		long:        "coulomb",
		value:       f64(1),
		dimension:   ELECTRIC_CHARGE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricPotentialUnits = []baseUnit{
	{
		name:        "V",
		long:        "volt",
		value:       f64(1),
		dimension:   ELECTRIC_POTENTIAL,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCapaticanceUnits = []baseUnit{
	{
		name:        "F",
		long:        "farad",
		value:       f64(1),
		dimension:   ELECTRIC_CAPACITANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricConductanceUnits = []baseUnit{
	{
		name:        "S",
		long:        "siemens",
		value:       f64(1),
		dimension:   ELECTRIC_CONDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxUnits = []baseUnit{
	{
		name:        "Wb",
		long:        "weber",
		value:       f64(1),
		dimension:   MAGNETIC_FLUX,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxDensityUnits = []baseUnit{
	{
		name:        "T",
		long:        "tesla",
		value:       f64(1),
		dimension:   MAGNETIC_FLUX_DENSITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricInductanceUnits = []baseUnit{
	{
		name:        "H",
		long:        "henry",
		value:       f64(1),
		dimension:   ELECTRIC_INDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricResistanceUnits = []baseUnit{
	{
		name:        "Ω, ohm",
		long:        "ohm",
		value:       f64(1),
		dimension:   ELECTRIC_RESISTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var solidAngleUnits = []baseUnit{
	{
		name:        "sr",
		long:        "steradian",
		value:       f64(1),
		dimension:   SOLID_ANGLE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var ionizingRadiationUnits = []baseUnit{
	{
		name:        "Sv",
		long:        "sievert",
		value:       f64(1),
		dimension:   IONIZING_RADIATION,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var radiationDoseUnits = []baseUnit{
	{
		name:        "Gy",
		long:        "gray",
		value:       f64(1),
		dimension:   RADIATION_DOSE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var catalycticActivityUnits = []baseUnit{
	{
		name:        "kat",
		long:        "katal",
		value:       f64(1),
		dimension:   CATALYCTIC_ACTIVITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var luminousFluxUnits = []baseUnit{
	{
		name:        "lm",
		long:        "lumen",
		value:       f64(1),
		dimension:   LUMINOUS_FLUX,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var illuminanceUnits = []baseUnit{
	{
		name:        "lx",
		long:        "lux",
		value:       f64(1),
		dimension:   ILLUMINANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}
