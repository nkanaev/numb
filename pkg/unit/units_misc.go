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
		name:      "a",
		aliases:   []string{"ampere"},
		value:     f64(1),
		dimension: ELECTRIC_CURRENT,
		prefixes:  &metricPrefixes,
	},
}

var luminousIntensityUnits = []baseUnit{
	{
		name:      "cd",
		aliases:   []string{"candela"},
		value:     f64(1),
		dimension: LUMINOUS_INTENSITY,
		prefixes:  &metricPrefixes,
	},
}

var amountOfSubstanceUnits = []baseUnit{
	{
		name:      "mol",
		aliases:   []string{"mole"},
		value:     f64(1),
		dimension: AMOUNT_OF_SUBSTANCE,
		prefixes:  &metricPrefixes,
	},
}
