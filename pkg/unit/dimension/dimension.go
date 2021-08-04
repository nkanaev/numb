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

	end_basis
)

type Dimensions [11]int

func (d1 Dimensions) Exp(x int) Dimensions {
	for b := Basis(0); b < end_basis; b++ {
		d1[b] *= x
	}
	return d1
}

func (d1 Dimensions) Add(d2 Dimensions) Dimensions {
	for b := Basis(0); b < end_basis; b++ {
		d1[b] += d2[b]
	}
	return d1
}

func (d1 Dimensions) Equals(d2 Dimensions) bool {
	for b := Basis(0); b < end_basis; b++ {
		if d1[b] != d2[b] {
			return false
		}
	}
	return true
}

func (d1 Dimensions) IsZero() bool {
	var nodim Dimensions
	return d1.Equals(nodim)
}

func (d1 Dimensions) Measure() Measure {
	for measure := MASS; measure < end_measures; measure++ {
		if measure.Dim().Equals(d1) {
			return measure
		}
	}
	return UNKNOWN
}
