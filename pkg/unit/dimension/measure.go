package dimension

type Measure uint

const (
	LENGTH Measure = iota
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

	end_measures
)

var measureDimensions = map[Measure]Dimensions{
	LENGTH:                dim{Length: 1}.Dim(),
	TEMPERATURE:           dim{Temperature: 1}.Dim(),
	AREA:                  dim{Length: 2}.Dim(),
	VOLUME:                dim{Length: 3}.Dim(),
	MASS:                  dim{Mass: 1}.Dim(),
	TIME:                  dim{Time: 1}.Dim(),
	ANGLE:                 dim{Angle: 1}.Dim(),
	DIGITAL:               dim{Digital: 1}.Dim(),
	CURRENCY:              dim{Currency: 1}.Dim(),
	FREQUENCY:             dim{Time: -1}.Dim(),
	ELECTRIC_CURRENT:      dim{Current: 1}.Dim(),
	LUMINOUS_INTENSITY:    dim{LuminousIntensity: 1}.Dim(),
	AMOUNT_OF_SUBSTANCE:   dim{AmountOfSubstance: 1}.Dim(),
	POWER:                 dim{Mass: 1, Length: 2, Time: -3}.Dim(),
	FORCE:                 dim{Mass: 1, Length: 1, Time: -2}.Dim(),
	ENERGY:                dim{Mass: 1, Length: 2, Time: -2}.Dim(),
	ELECTRIC_CHARGE:       dim{Time: 1, Current: 1}.Dim(),
	ELECTRIC_POTENTIAL:    dim{Mass: 1, Length: 2, Time: -3, Current: -1}.Dim(),
	ELECTRIC_CAPACITANCE:  dim{Mass: -1, Length: -2, Time: 4, Current: 2}.Dim(),
	ELECTRIC_CONDUCTANCE:  dim{Mass: -1, Length: -2, Time: 3, Current: 2}.Dim(),
	MAGNETIC_FLUX:         dim{Mass: 1, Length: 2, Time: -2, Current: -1}.Dim(),
	MAGNETIC_FLUX_DENSITY: dim{Mass: 1, Time: -2, Current: -1}.Dim(),
	ELECTRIC_INDUCTANCE:   dim{Mass: 1, Length: 2, Time: -2, Current: -2}.Dim(),
	ELECTRIC_RESISTANCE:   dim{Mass: 1, Length: 2, Time: -3, Current: 2}.Dim(),
	PRESSURE:              dim{Mass: 1, Length: -1, Time: -2}.Dim(),
	RADIOACTIVITY:         dim{Time: -1}.Dim(),
	SOLID_ANGLE:           dim{SolidAngle: 1}.Dim(),
	IONIZING_RADIATION:    dim{Length: 2, Time: -2}.Dim(),
	CATALYCTIC_ACTIVITY:   dim{AmountOfSubstance: 1, Time: -1}.Dim(),
	RADIATION_DOSE:        dim{Length: 2, Time: -2}.Dim(),
	LUMINOUS_FLUX:         dim{LuminousIntensity: 1, SolidAngle: 1}.Dim(),
	ILLUMINANCE:           dim{LuminousIntensity: 1, SolidAngle: 1, Length: -1}.Dim(),
}

var measureNames = map[Measure]string{
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
}

func (d Measure) Dim() Dimensions {
	return measureDimensions[d]
}

func (d Measure) String() string {
	return measureNames[d]
}
