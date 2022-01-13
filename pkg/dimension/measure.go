package dimension

type Measure int

const (
	start_measures Measure = iota
	DIMENSIONLESS
	LENGTH
	TEMPERATURE
	AREA
	VOLUME
	MASS
	TIME
	ANGLE
	DIGITAL
	CURRENCY
	FREQUENCY
	ELECTRIC_CURRENT
	LUMINOUS_INTENSITY
	AMOUNT_OF_SUBSTANCE
	POWER
	FORCE
	ENERGY
	ELECTRIC_CHARGE
	ELECTRIC_POTENTIAL
	ELECTRIC_CAPACITANCE
	ELECTRIC_CONDUCTANCE
	MAGNETIC_FLUX
	MAGNETIC_FLUX_DENSITY
	ELECTRIC_INDUCTANCE
	ELECTRIC_RESISTANCE
	PRESSURE
	RADIOACTIVITY
	SOLID_ANGLE
	IONIZING_RADIATION
	CATALYCTIC_ACTIVITY
	RADIATION_DOSE
	LUMINOUS_FLUX
	ILLUMINANCE
	SPEED
	DATA_RATE
	DENSITY
	FLOW
	ACCELERATION

	end_measures
)

func List() []Measure {
	result := make([]Measure, end_measures - start_measures)
	for measure := start_measures+1; measure < end_measures; measure++ {
		result = append(result, measure)
	}
	return result
}

var measureDimensions = map[Measure]Dimensions{
	DIMENSIONLESS:         Dimensions{},
	LENGTH:                Dimensions{Length: 1},
	TEMPERATURE:           Dimensions{Temperature: 1},
	AREA:                  Dimensions{Length: 2},
	VOLUME:                Dimensions{Length: 3},
	MASS:                  Dimensions{Mass: 1},
	TIME:                  Dimensions{Time: 1},
	ANGLE:                 Dimensions{},
	DIGITAL:               Dimensions{Digital: 1},
	CURRENCY:              Dimensions{Currency: 1},
	FREQUENCY:             Dimensions{Time: -1},
	ELECTRIC_CURRENT:      Dimensions{Current: 1},
	LUMINOUS_INTENSITY:    Dimensions{LuminousIntensity: 1},
	AMOUNT_OF_SUBSTANCE:   Dimensions{AmountOfSubstance: 1},
	POWER:                 Dimensions{Mass: 1, Length: 2, Time: -3},
	FORCE:                 Dimensions{Mass: 1, Length: 1, Time: -2},
	ENERGY:                Dimensions{Mass: 1, Length: 2, Time: -2},
	ELECTRIC_CHARGE:       Dimensions{Time: 1, Current: 1},
	ELECTRIC_POTENTIAL:    Dimensions{Mass: 1, Length: 2, Time: -3, Current: -1},
	ELECTRIC_CAPACITANCE:  Dimensions{Mass: -1, Length: -2, Time: 4, Current: 2},
	ELECTRIC_CONDUCTANCE:  Dimensions{Mass: -1, Length: -2, Time: 3, Current: 2},
	MAGNETIC_FLUX:         Dimensions{Mass: 1, Length: 2, Time: -2, Current: -1},
	MAGNETIC_FLUX_DENSITY: Dimensions{Mass: 1, Time: -2, Current: -1},
	ELECTRIC_INDUCTANCE:   Dimensions{Mass: 1, Length: 2, Time: -2, Current: -2},
	ELECTRIC_RESISTANCE:   Dimensions{Mass: 1, Length: 2, Time: -3, Current: 2},
	PRESSURE:              Dimensions{Mass: 1, Length: -1, Time: -2},
	RADIOACTIVITY:         Dimensions{Time: -1},
	SOLID_ANGLE:           Dimensions{},
	IONIZING_RADIATION:    Dimensions{Length: 2, Time: -2},
	CATALYCTIC_ACTIVITY:   Dimensions{AmountOfSubstance: 1, Time: -1},
	RADIATION_DOSE:        Dimensions{Length: 2, Time: -2},
	LUMINOUS_FLUX:         Dimensions{LuminousIntensity: 1, SolidAngle: 1},
	ILLUMINANCE:           Dimensions{LuminousIntensity: 1, SolidAngle: 1, Length: -1},
	SPEED:                 Dimensions{Length: 1, Time: -1},
	DATA_RATE:             Dimensions{Digital: 1, Time: -1},
	DENSITY:               Dimensions{Mass: 1, Length: -3},
	FLOW:                  Dimensions{Length: 3, Time: -1},
	ACCELERATION:          Dimensions{Length: 1, Time: -2},
}

var measureNames = map[Measure]string{
	DIMENSIONLESS:         "DIMENSIONLESS",
	LENGTH:                "LENGTH",
	TEMPERATURE:           "TEMPERATURE",
	AREA:                  "AREA",
	VOLUME:                "VOLUME",
	MASS:                  "MASS",
	TIME:                  "TIME",
	ANGLE:                 "ANGLE",
	DIGITAL:               "DIGITAL",
	CURRENCY:              "CURRENCY",
	FREQUENCY:             "FREQUENCY",
	ELECTRIC_CURRENT:      "ELECTRIC_CURRENT",
	LUMINOUS_INTENSITY:    "LUMINOUS_INTENSITY",
	AMOUNT_OF_SUBSTANCE:   "AMOUNT_OF_SUBSTANCE",
	POWER:                 "POWER",
	FORCE:                 "FORCE",
	ENERGY:                "ENERGY",
	ELECTRIC_CHARGE:       "ELECTRIC_CHARGE",
	ELECTRIC_POTENTIAL:    "ELECTRIC_POTENTIAL",
	ELECTRIC_CAPACITANCE:  "ELECTRIC_CAPACITANCE",
	ELECTRIC_CONDUCTANCE:  "ELECTRIC_CONDUCTANCE",
	MAGNETIC_FLUX:         "MAGNETIC_FLUX",
	MAGNETIC_FLUX_DENSITY: "MAGNETIC_FLUX_DENSITY",
	ELECTRIC_INDUCTANCE:   "ELECTRIC_INDUCTANCE",
	ELECTRIC_RESISTANCE:   "ELECTRIC_RESISTANCE",
	PRESSURE:              "PRESSURE",
	RADIOACTIVITY:         "RADIOACTIVITY",
	SOLID_ANGLE:           "SOLID_ANGLE",
	IONIZING_RADIATION:    "IONIZING_RADIATION",
	CATALYCTIC_ACTIVITY:   "CATALYCTIC_ACTIVITY",
	RADIATION_DOSE:        "RADIATION_DOSE",
	LUMINOUS_FLUX:         "LUMINOUS_FLUX",
	ILLUMINANCE:           "ILLUMINANCE",
	SPEED:                 "SPEED",
	DATA_RATE:             "DATA_RANGE",
	DENSITY:               "DENSITY",
	FLOW:                  "FLOW",
	ACCELERATION:          "ACCELERATION",
}

func (d Measure) Dim() Dimensions {
	return measureDimensions[d]
}

func (d Measure) String() string {
	name, found := measureNames[d]
	if !found {
		return "unknown"
	}
	return name
}
