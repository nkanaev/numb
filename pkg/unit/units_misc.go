package unit

var frequencyUnits = []baseUnit{
	{
		short:       "Hz",
		long:        "hertz",
		value:       f64(1),
		dimension:   FREQUENCY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCurrentUnits = []baseUnit{
	{
		short:       "A, amp",
		long:        "ampere",
		value:       f64(1),
		dimension:   ELECTRIC_CURRENT,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var luminousIntensityUnits = []baseUnit{
	{
		short:       "cd",
		long:        "candela",
		value:       f64(1),
		dimension:   LUMINOUS_INTENSITY,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var amountOfSubstanceUnits = []baseUnit{
	{
		short:       "mol",
		long:        "mole",
		value:       f64(1),
		dimension:   AMOUNT_OF_SUBSTANCE,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var electricChargeUnits = []baseUnit{
	{
		short:       "C",
		long:        "coulomb",
		value:       f64(1),
		dimension:   ELECTRIC_CHARGE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricPotentialUnits = []baseUnit{
	{
		short:       "V",
		long:        "volt",
		value:       f64(1),
		dimension:   ELECTRIC_POTENTIAL,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCapaticanceUnits = []baseUnit{
	{
		short:       "F",
		long:        "farad",
		value:       f64(1),
		dimension:   ELECTRIC_CAPACITANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricConductanceUnits = []baseUnit{
	{
		short:       "S",
		long:        "siemens",
		value:       f64(1),
		dimension:   ELECTRIC_CONDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxUnits = []baseUnit{
	{
		short:       "Wb",
		long:        "weber",
		value:       f64(1),
		dimension:   MAGNETIC_FLUX,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxDensityUnits = []baseUnit{
	{
		short:       "T",
		long:        "tesla",
		value:       f64(1),
		dimension:   MAGNETIC_FLUX_DENSITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricInductanceUnits = []baseUnit{
	{
		short:       "H",
		long:        "henry",
		value:       f64(1),
		dimension:   ELECTRIC_INDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricResistanceUnits = []baseUnit{
	{
		short:       "Ω, ohm",
		long:        "ohm",
		value:       f64(1),
		dimension:   ELECTRIC_RESISTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var solidAngleUnits = []baseUnit{
	{
		short:       "sr",
		long:        "steradian",
		value:       f64(1),
		dimension:   SOLID_ANGLE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var ionizingRadiationUnits = []baseUnit{
	{
		short:       "Sv",
		long:        "sievert",
		value:       f64(1),
		dimension:   IONIZING_RADIATION,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var radiationDoseUnits = []baseUnit{
	{
		short:       "Gy",
		long:        "gray",
		value:       f64(1),
		dimension:   RADIATION_DOSE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var catalycticActivityUnits = []baseUnit{
	{
		short:       "kat",
		long:        "katal",
		value:       f64(1),
		dimension:   CATALYCTIC_ACTIVITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var luminousFluxUnits = []baseUnit{
	{
		short:       "lm",
		long:        "lumen",
		value:       f64(1),
		dimension:   LUMINOUS_FLUX,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var illuminanceUnits = []baseUnit{
	{
		short:       "lx",
		long:        "lux",
		value:       f64(1),
		dimension:   ILLUMINANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}
