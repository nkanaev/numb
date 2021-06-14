package unit

type dim struct {
	Name                           string
	Mass, Length, Time, Current    int
	Temperature, LuminousIntensity int
	AmountOfSubstance              int
	Angle, SolidAngle              int
	Digital, Currency              int
}

type BaseUnit *dim

var (
	LENGTH                = &dim{Name: "length", Length: 1}
	TEMPERATURE           = &dim{Name: "temperature", Temperature: 1}
	AREA                  = &dim{Name: "area", Length: 2}
	VOLUME                = &dim{Name: "volume", Length: 3}
	MASS                  = &dim{Name: "mass", Mass: 1}
	TIME                  = &dim{Name: "time", Time: 1}
	ANGLE                 = &dim{Name: "angle", Angle: 1}
	DIGITAL               = &dim{Name: "digital", Digital: 1}
	CURRENCY              = &dim{Name: "currency", Currency: 1}
	FREQUENCY             = &dim{Name: "frequency", Time: -1}
	ELECTRIC_CURRENT      = &dim{Name: "electric current", Current: 1}
	LUMINOUS_INTENSITY    = &dim{Name: "luminous intensity", LuminousIntensity: 1}
	AMOUNT_OF_SUBSTANCE   = &dim{Name: "amount of substance", AmountOfSubstance: 1}
	POWER                 = &dim{Name: "power", Mass: 1, Length: 2, Time: -3}
	FORCE                 = &dim{Name: "force", Mass: 1, Length: 1, Time: -2}
	ENERGY                = &dim{Name: "energy", Mass: 1, Length: 2, Time: -2}
	ELECTRIC_CHARGE       = &dim{Name: "electric charge", Time: 1, Current: 1}
	ELECTRIC_POTENTIAL    = &dim{Name: "electric potential", Mass: 1, Length: 2, Time: -3, Current: -1}
	ELECTRIC_CAPACITANCE  = &dim{Name: "electric capacitance", Mass: -1, Length: -2, Time: 4, Current: 2}
	ELECTRIC_CONDUCTANCE  = &dim{Name: "electric conductaance", Mass: -1, Length: -2, Time: 3, Current: 2}
	MAGNETIC_FLUX         = &dim{Name: "magnetic flux", Mass: 1, Length: 2, Time: -2, Current: -1}
	MAGNETIC_FLUX_DENSITY = &dim{Name: "magnetic flux density", Mass: 1, Time: -2, Current: -1}
	ELECTRIC_INDUCTANCE   = &dim{Name: "electric inductance", Mass: 1, Length: 2, Time: -2, Current: -2}
	ELECTRIC_RESISTANCE   = &dim{Name: "electric resistance", Mass: 1, Length: 2, Time: -3, Current: 2}
	PRESSURE              = &dim{Name: "pressure", Mass: 1, Length: -1, Time: -2}
	RADIOACTIVITY         = &dim{Name: "radioactivity", Time: -1}
	SOLID_ANGLE           = &dim{Name: "solid angle", SolidAngle: 1}
	IONIZING_RADIATION    = &dim{Name: "ionizing radiation", Length: 2, Time: -2}
	CATALYCTIC_ACTIVITY   = &dim{Name: "catalyctic activity", AmountOfSubstance: 1, Time: -1}
	RADIATION_DOSE        = &dim{Name: "radiation dose", Length: 2, Time: -2}
	LUMINOUS_FLUX         = &dim{Name: "luminous flux", LuminousIntensity: 1, SolidAngle: 1}
	ILLUMINANCE           = &dim{Name: "illuminance", LuminousIntensity: 1, SolidAngle: 1, Length: -1}
)
