package unit

var frequencyUnits = []baseUnit{
	{
		name:        "Hz",
		aliases:     []string{"hertz"},
		value:       f64(1),
		dimension:   FREQUENCY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCurrentUnits = []baseUnit{
	{
		name:         "A",
		shortaliases: []string{"amp"},
		aliases:      []string{"ampere"},
		value:        f64(1),
		dimension:    ELECTRIC_CURRENT,
		prefixes:     &metricPrefixes,
		description:  "SI base unit",
	},
}

var luminousIntensityUnits = []baseUnit{
	{
		name:        "cd",
		aliases:     []string{"candela"},
		value:       f64(1),
		dimension:   LUMINOUS_INTENSITY,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var amountOfSubstanceUnits = []baseUnit{
	{
		name:        "mol",
		aliases:     []string{"mole"},
		value:       f64(1),
		dimension:   AMOUNT_OF_SUBSTANCE,
		prefixes:    &metricPrefixes,
		description: "SI base unit",
	},
}

var electricChargeUnits = []baseUnit{
	{
		name:        "C",
		aliases:     []string{"coulomb"},
		value:       f64(1),
		dimension:   ELECTRIC_CHARGE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricPotentialUnits = []baseUnit{
	{
		name:        "V",
		aliases:     []string{"volt"},
		value:       f64(1),
		dimension:   ELECTRIC_POTENTIAL,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricCapaticanceUnits = []baseUnit{
	{
		name:        "F",
		aliases:     []string{"farad"},
		value:       f64(1),
		dimension:   ELECTRIC_CAPACITANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricConductanceUnits = []baseUnit{
	{
		name:        "S",
		aliases:     []string{"siemens"},
		value:       f64(1),
		dimension:   ELECTRIC_CONDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxUnits = []baseUnit{
	{
		name:        "Wb",
		aliases:     []string{"weber"},
		value:       f64(1),
		dimension:   MAGNETIC_FLUX,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var magneticFluxDensityUnits = []baseUnit{
	{
		name:        "T",
		aliases:     []string{"tesla"},
		value:       f64(1),
		dimension:   MAGNETIC_FLUX_DENSITY,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricInductanceUnits = []baseUnit{
	{
		name:        "H",
		aliases:     []string{"henry"},
		value:       f64(1),
		dimension:   ELECTRIC_INDUCTANCE,
		prefixes:    &metricPrefixes,
		description: "SI derived unit",
	},
}

var electricResistanceUnits = []baseUnit{
	{
		name:         "Ω",
		shortaliases: []string{"ohm"},
		aliases:      []string{"ohm"},
		value:        f64(1),
		dimension:    ELECTRIC_RESISTANCE,
		prefixes:     &metricPrefixes,
		description:  "SI derived unit",
	},
}

var solidAngleUnits = []baseUnit{
	{
		name:         "sr",
		aliases:      []string{"steradian"},
		value:        f64(1),
		dimension:    SOLID_ANGLE,
		prefixes:     &metricPrefixes,
		description:  "SI derived unit",
	},
}

var ionizingRadiationUnits = []baseUnit{
	{
		name:         "Sv",
		aliases:      []string{"sievert"},
		value:        f64(1),
		dimension:    IONIZING_RADIATION,
		prefixes:     &metricPrefixes,
		description:  "SI derived unit",
	},
}
