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

	"cosh":  nop,
	"sinh":  nop,
	"tanh":  nop,
	"acosh": nop,
	"asinh": nop,
	"atanh": nop,

	"ln":    nop,
	"log":   nop,
	"log2":  nop,
	"log10": nop,
	"sqrt":  nop,

	"ceil":  nop,
	"floor": nop,
	"abs":   Abs,
	"trunc": nop,

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
	return value.NewInt(1)
}
