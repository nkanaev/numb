package dimension

type Basis uint

const (
	Mass Basis = iota
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

type Dimension struct {
	Name string
	Dims [11]int
}

func (d Dimension) String() string {
	return d.Name
}

type dim map[Basis]int

func (d dim) Dim(name string) Dimension {
	var r Dimension
	for basis, exp := range d {
		r.Dims[basis] = exp
	}
	r.Name = name
	return r
}

var (
	LENGTH                = dim{Length: 1}.Dim("LENGTH")
	TEMPERATURE           = dim{Temperature: 1}.Dim("TEMPERATURE")
	AREA                  = dim{Length: 2}.Dim("AREA")
	VOLUME                = dim{Length: 3}.Dim("VOLUME")
	MASS                  = dim{Mass: 1}.Dim("MASS")
	TIME                  = dim{Time: 1}.Dim("TIME")
	ANGLE                 = dim{Angle: 1}.Dim("ANGLE")
	DIGITAL               = dim{Digital: 1}.Dim("DIGITAL")
	CURRENCY              = dim{Currency: 1}.Dim("CURRENCY")
	FREQUENCY             = dim{Time: -1}.Dim("FREQUENCY")
	ELECTRIC_CURRENT      = dim{Current: 1}.Dim("ELECTRIC_CURRENT")
	LUMINOUS_INTENSITY    = dim{LuminousIntensity: 1}.Dim("LUMINOUS_INTENSITY")
	AMOUNT_OF_SUBSTANCE   = dim{AmountOfSubstance: 1}.Dim("AMOUNT_OF_SUBSTANCE")
	POWER                 = dim{Mass: 1, Length: 2, Time: -3}.Dim("POWER")
	FORCE                 = dim{Mass: 1, Length: 1, Time: -2}.Dim("FORCE")
	ENERGY                = dim{Mass: 1, Length: 2, Time: -2}.Dim("ENERGY")
	ELECTRIC_CHARGE       = dim{Time: 1, Current: 1}.Dim("ELECTRIC_CHARGE")
	ELECTRIC_POTENTIAL    = dim{Mass: 1, Length: 2, Time: -3, Current: -1}.Dim("ELECTRIC_POTENTIAL")
	ELECTRIC_CAPACITANCE  = dim{Mass: -1, Length: -2, Time: 4, Current: 2}.Dim("ELECTRIC_CAPACITANCE")
	ELECTRIC_CONDUCTANCE  = dim{Mass: -1, Length: -2, Time: 3, Current: 2}.Dim("ELECTRIC_CONDUCTANCE")
	MAGNETIC_FLUX         = dim{Mass: 1, Length: 2, Time: -2, Current: -1}.Dim("MAGNETIC_FLUX")
	MAGNETIC_FLUX_DENSITY = dim{Mass: 1, Time: -2, Current: -1}.Dim("MAGNETIC_FLUX_DENSITY")
	ELECTRIC_INDUCTANCE   = dim{Mass: 1, Length: 2, Time: -2, Current: -2}.Dim("ELECTRIC_INDUCTANCE")
	ELECTRIC_RESISTANCE   = dim{Mass: 1, Length: 2, Time: -3, Current: 2}.Dim("ELECTRIC_RESISTANCE")
	PRESSURE              = dim{Mass: 1, Length: -1, Time: -2}.Dim("PRESSURE")
	RADIOACTIVITY         = dim{Time: -1}.Dim("RADIOACTIVITY")
	SOLID_ANGLE           = dim{SolidAngle: 1}.Dim("SOLID_ANGLE")
	IONIZING_RADIATION    = dim{Length: 2, Time: -2}.Dim("IONIZING_RADIATION")
	CATALYCTIC_ACTIVITY   = dim{AmountOfSubstance: 1, Time: -1}.Dim("CATALYCTIC_ACTIVITY")
	RADIATION_DOSE        = dim{Length: 2, Time: -2}.Dim("RADIATION_DOSE")
	LUMINOUS_FLUX         = dim{LuminousIntensity: 1, SolidAngle: 1}.Dim("LUMINOUS_FLUX")
	ILLUMINANCE           = dim{LuminousIntensity: 1, SolidAngle: 1, Length: -1}.Dim("ILLUMINANCE")
)
