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
	Digital
	Currency

	end_basis
)

type Dimensions [9]int

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

func (d1 Dimensions) Measure() string {
	for name, dim := range Measures {
		if dim == d1 {
			return name
		}
	}
	return "unknown"
}

var (
	LENGTH              = Dimensions{Length: 1}
	TIME                = Dimensions{Time: 1}
	MASS                = Dimensions{Mass: 1}
	ELECTRIC_CURRENT    = Dimensions{Current: 1}
	TEMPERATURE         = Dimensions{Temperature: 1}
	LUMINOUS_INTENSITY  = Dimensions{LuminousIntensity: 1}
	AMOUNT_OF_SUBSTANCE = Dimensions{AmountOfSubstance: 1}
	DIGITAL             = Dimensions{Digital: 1}
	CURRENCY            = Dimensions{Currency: 1}
)
