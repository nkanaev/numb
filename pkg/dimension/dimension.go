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

type Dimension [9]int

func (d1 Dimension) Exp(x int) Dimension {
	for b := Basis(0); b < end_basis; b++ {
		d1[b] *= x
	}
	return d1
}

func (d1 Dimension) Add(d2 Dimension) Dimension {
	for b := Basis(0); b < end_basis; b++ {
		d1[b] += d2[b]
	}
	return d1
}

func (d1 Dimension) Equals(d2 Dimension) bool {
	for b := Basis(0); b < end_basis; b++ {
		if d1[b] != d2[b] {
			return false
		}
	}
	return true
}

func (d1 Dimension) IsZero() bool {
	var nodim Dimension
	return d1.Equals(nodim)
}

func (d1 Dimension) Measure() string {
	for name, dim := range Measures {
		if dim == d1 {
			return name
		}
	}
	return "unknown"
}

var (
	LENGTH              = Dimension{Length: 1}
	TIME                = Dimension{Time: 1}
	MASS                = Dimension{Mass: 1}
	ELECTRIC_CURRENT    = Dimension{Current: 1}
	TEMPERATURE         = Dimension{Temperature: 1}
	LUMINOUS_INTENSITY  = Dimension{LuminousIntensity: 1}
	AMOUNT_OF_SUBSTANCE = Dimension{AmountOfSubstance: 1}
	DIGITAL             = Dimension{Digital: 1}
	CURRENCY            = Dimension{Currency: 1}
)
