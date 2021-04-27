package funcs

import (
	"math"
	"math/big"

	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

func Sin(args ...value.Value) value.Value {
	arg := args[0]
	if arg.Unit == nil {
		panic("provide rad or deg unit")
	}
	u := arg.Unit.String()
	if u != "deg" && u != "rad" {
		panic("expected rad or deg unit")
	}
	if u == "deg" {
		arg = arg.To(unit.Get("rad"))
	}
	f64, _ := arg.Num.Float64()
	num := new(big.Rat).SetFloat64(math.Sin(f64))
	return value.Value{Num: num}
}
