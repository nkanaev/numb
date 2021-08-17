package value

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/ratutils"
)

type Format int

const (
	DEC Format = iota
	HEX
	OCT
	BIN
	RAT
	SCI
)

var FormatToString = map[Format]string{
	DEC: "dec",
	HEX: "hex",
	OCT: "oct",
	BIN: "bin",
	RAT: "rat",
	SCI: "sci",
}

var StringToFormat = map[string]Format{}

func (n Format) String() string {
	return FormatToString[n]
}

func (a Value) String() string {
	return a.Format(",", 2)
}

func (a Value) Format(sep string, prec int) string {
	num := ""
	switch a.Fmt {
	case DEC:
		num = formatDec(a.Num, sep, prec)
	case HEX:
		num = fmt.Sprintf("%#x", ratutils.ToInt(a.Num))
	case OCT:
		num = fmt.Sprintf("%O", ratutils.ToInt(a.Num))
	case BIN:
		num = fmt.Sprintf("%#b", ratutils.ToInt(a.Num))
	case RAT:
		num = fmt.Sprintf(
			"%s/%s",
			groupDigits(a.Num.Num().String(), sep, 3),
			groupDigits(a.Num.Denom().String(), sep, 3),
		)
	case SCI:
		num = formatSci(a.Num, sep, prec)
	}
	if a.Unit != nil {
		num += " " + a.Unit.String()
	}
	return num
}

func groupDigits(num string, sep string, size int) string {
	if len(sep) == 0 {
		return num
	}
	ch := rune(sep[0])
	digits := []rune(num)
	out := make([]rune, 0)
	for i := 0; i < len(digits); i++ {
		ord := len(digits) - i
		if ord != len(num) && (ord % size) == 0 {
			out = append(out, ch)
		}
		out = append(out, digits[i])
	}	
	return string(out)
}

func formatSci(rat *big.Rat, sep string, prec int) string {
	str := fmt.Sprintf(fmt.Sprint("%.",  prec, "e"), new(big.Float).SetRat(rat))
	
	var int, dec, exp, expsign string

	parts := strings.Split(str, ".")
	int = parts[0]
	dec = parts[1]

	parts = strings.Split(dec, "e")
	dec = parts[0]
	exp = parts[1]

	expsign = exp[:1]
	if expsign == "+" {
		expsign = ""
	}
	exp = exp[1:]
	dec = strings.TrimRight(dec, "0")
	exp = strings.TrimLeft(exp, "0")
	if len(exp) == 0 {
		exp = "0"
	}

	out := int
	if len(dec) > 0 {
		out += "." + dec
	}
	out += "e" + expsign + exp
	return out
}

func formatDec(rat *big.Rat, sep string, prec int) string {
	var str string
	if rat.IsInt() {
		str = rat.RatString()
	} else {
		str = rat.FloatString(prec)
	}

	var sign, int, dec string

	if str[0] == '-' {
		sign = "-"
		str = str[1:]
	}

	parts := strings.Split(str, ".")
	int = parts[0]
	if len(parts) == 2 && prec > 0 {
		dec = parts[1]
	}

	var out string
	out = sign + groupDigits(int, sep, 3)
	if len(dec) > 0 {
		dec = strings.TrimRight(dec, "0")
		if len(dec) == 0 {
			dec = "0"
		}
		out += "." + dec
	}
	return out
}

func init() {
	for num, str := range FormatToString {
		StringToFormat[str] = num
	}
}
