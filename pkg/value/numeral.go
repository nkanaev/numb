package value

type NumeralSystem int

const (
	DEC NumeralSystem = iota
	HEX
	OCT
	BIN
	RAT
	SCI
)

var NumeralToString = map[NumeralSystem]string{
	DEC: "dec",
	HEX: "hex",
	OCT: "oct",
	BIN: "bin",
	RAT: "rat",
	SCI: "sci",
}

var StringToNumeral = map[string]NumeralSystem{}

func (n NumeralSystem) String() string {
	return NumeralToString[n]
}

func init() {
	for num, str := range NumeralToString {
		StringToNumeral[str] = num
	}
}
