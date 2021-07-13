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
	end_dimensions
)

type Dimensions [11]int

func (d1 Dimensions) Exp(x int) Dimensions {
	for b := Mass; b < end_dimensions; b++ {
		d1[b] *= x
	}
	return d1
}

func (d1 Dimensions) Add(d2 Dimensions) Dimensions {
	for b := Mass; b < end_dimensions; b++ {
		d1[b] += d2[b]
	}
	return d1
}

func (d1 Dimensions) Equals(d2 Dimensions) bool {
	for b := Mass; b < end_dimensions; b++ {
		if d1[b] != d2[b] {
			return false
		}
	}
	return true
}
