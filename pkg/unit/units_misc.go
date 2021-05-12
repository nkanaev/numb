package unit

var frequencyUnits = []baseUnit{
	{
		name:      "hz",
		aliases:   []string{"hertz"},
		value:     f64(1),
		dimension: FREQUENCY,
		prefixes:  &metricPrefixes,
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
