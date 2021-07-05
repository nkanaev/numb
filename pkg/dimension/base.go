package dimension

type BaseDimension uint

const (
	Mass BaseDimension = iota
	Length
	Time
	Current
	Temperature
	LuminousIntensity
	AmountOfSubstance
	Angle
	SolidAngle
	Digital
	Currency
)

type Dimension [11]int

type dim map[BaseDimension]int

func (d dim) Dim() Dimension {
	var r Dimension
	for dim, exp := range d {
		d[dim] = exp
	}
	return r
}

var (
	LENGTH                = dim{Length: 1}.Dim()
	TEMPERATURE           = dim{Temperature: 1}.Dim()
	AREA                  = dim{Length: 2}.Dim()
	VOLUME                = dim{Length: 3}.Dim()
	MASS                  = dim{Mass: 1}.Dim()
	TIME                  = dim{Time: 1}.Dim()
	ANGLE                 = dim{Angle: 1}.Dim()
	DIGITAL               = dim{Digital: 1}.Dim()
	CURRENCY              = dim{Currency: 1}.Dim()
	FREQUENCY             = dim{Time: -1}.Dim()
	ELECTRIC_CURRENT      = dim{Current: 1}.Dim()
	LUMINOUS_INTENSITY    = dim{LuminousIntensity: 1}.Dim()
	AMOUNT_OF_SUBSTANCE   = dim{AmountOfSubstance: 1}.Dim()
	POWER                 = dim{Mass: 1, Length: 2, Time: -3}.Dim()
	FORCE                 = dim{Mass: 1, Length: 1, Time: -2}.Dim()
	ENERGY                = dim{Mass: 1, Length: 2, Time: -2}.Dim()
	ELECTRIC_CHARGE       = dim{Time: 1, Current: 1}.Dim()
	ELECTRIC_POTENTIAL    = dim{Mass: 1, Length: 2, Time: -3, Current: -1}.Dim()
	ELECTRIC_CAPACITANCE  = dim{Mass: -1, Length: -2, Time: 4, Current: 2}.Dim()
	ELECTRIC_CONDUCTANCE  = dim{Mass: -1, Length: -2, Time: 3, Current: 2}.Dim()
	MAGNETIC_FLUX         = dim{Mass: 1, Length: 2, Time: -2, Current: -1}.Dim()
	MAGNETIC_FLUX_DENSITY = dim{Mass: 1, Time: -2, Current: -1}.Dim()
	ELECTRIC_INDUCTANCE   = dim{Mass: 1, Length: 2, Time: -2, Current: -2}.Dim()
	ELECTRIC_RESISTANCE   = dim{Mass: 1, Length: 2, Time: -3, Current: 2}.Dim()
	PRESSURE              = dim{Mass: 1, Length: -1, Time: -2}.Dim()
	RADIOACTIVITY         = dim{Time: -1}.Dim()
	SOLID_ANGLE           = dim{SolidAngle: 1}.Dim()
	IONIZING_RADIATION    = dim{Length: 2, Time: -2}.Dim()
	CATALYCTIC_ACTIVITY   = dim{AmountOfSubstance: 1, Time: -1}.Dim()
	RADIATION_DOSE        = dim{Length: 2, Time: -2}.Dim()
	LUMINOUS_FLUX         = dim{LuminousIntensity: 1, SolidAngle: 1}.Dim()
	ILLUMINANCE           = dim{LuminousIntensity: 1, SolidAngle: 1, Length: -1}.Dim()
)
