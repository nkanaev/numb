package funcs

import "github.com/nkanaev/numb/pkg/value"

type Func func(...value.Value) value.Value

var db = map[string]Func{
	"sin":  Sin,
	"cos":  Cos,
	"tan":  Tan,
	"asin": Asin,
	"acos": Acos,
	"atan": Atan,

	"log":   Log,
	"log2":  Log2,
	"log10": Log10,
	"sqrt":  Sqrt,

	"ceil":  Ceil,
	"floor": Floor,
	"trunc": Trunc,
	"round": nop,
	"abs":   Abs,

	"gcd": GCD,
	"lcm": LCM,
}

func Get(x string) *Func {
	if fn, ok := db[x]; ok {
		return &fn
	}
	return nil
}

func nop(args ...value.Value) value.Value {
	return value.Int64(1)
}
