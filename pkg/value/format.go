package value

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/ratutils"
)

type NumberFormat int

const (
	DEC NumberFormat = iota
	HEX
	OCT
	BIN
	RAT
	SCI
)

var FormatToString = map[NumberFormat]string{
	DEC: "dec",
	HEX: "hex",
	OCT: "oct",
	BIN: "bin",
	RAT: "rat",
	SCI: "sci",
}

var StringToFormat = map[string]NumberFormat{}

func init() {
	for num, str := range FormatToString {
		StringToFormat[str] = num
	}
}

func (n NumberFormat) String() string {
	return FormatToString[n]
}

func formatNum(num *big.Rat, numfmt NumberFormat, sep string, prec int) string {
	ret := ""
	switch numfmt {
	case DEC:
		ret = formatDec(num, sep, prec)
	case HEX:
		ret = fmt.Sprintf("%#x", ratutils.TruncInt(num))
	case OCT:
		ret = fmt.Sprintf("%O", ratutils.TruncInt(num))
	case BIN:
		ret = fmt.Sprintf("%#b", ratutils.TruncInt(num))
	case RAT:
		ret = fmt.Sprintf(
			"%s/%s",
			groupDigits(num.Num().String(), sep, 3),
			groupDigits(num.Denom().String(), sep, 3),
		)
	case SCI:
		ret = formatSci(num, sep, prec)
	}
	return ret
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
		if ord != len(num) && (ord%size) == 0 {
			out = append(out, ch)
		}
		out = append(out, digits[i])
	}
	return string(out)
}

func formatSci(rat *big.Rat, sep string, prec int) string {
	str := fmt.Sprintf(fmt.Sprint("%.", prec, "e"), new(big.Float).SetRat(rat))

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
