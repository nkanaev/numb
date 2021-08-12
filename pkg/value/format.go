package value

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/ratutils"
)

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
		num = fmt.Sprintf(fmt.Sprint("%.",  prec, "e"), new(big.Float).SetRat(a.Num))
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

func formatDec(rat *big.Rat, sep string, prec int) string {
	var str string
	if rat.IsInt() {
		str = rat.RatString()
	} else {
		str = rat.FloatString(prec)
	}

	var sign, num, dec string

	if str[0] == '-' {
		sign = "-"
		str = str[1:]
	}

	parts := strings.Split(str, ".")
	num = parts[0]
	if len(parts) == 2 && prec > 0 {
		dec = parts[1]
	}

	var out string
	out = sign + groupDigits(num, sep, 3)
	if len(dec) > 0 {
		dec = strings.TrimRight(dec, "0")
		if len(dec) == 0 {
			dec = "0"
		}
		out += "." + dec
	}
	return out
}
