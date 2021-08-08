package funcs

import (
	"fmt"
	"math"

	"github.com/nkanaev/numb/pkg/value"
)

func arity(want, have int) {
	if want != have {
		panic(fmt.Sprintf("expected %d arguments, got %d", want, have))
	}
}

func f64(val value.Value) float64 {
	f, exact := val.Num.Float64()	
	if !exact && math.IsInf(f, 0) {
		panic(fmt.Sprintf("%s magnitude is too large", val.String()))
	}
	return f
}

func Log(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.FromFloat64(math.Log(f64(args[0])))
}

func Log2(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.FromFloat64(math.Log2(f64(args[0])))
}

func Log10(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.FromFloat64(math.Log10(f64(args[0])))
}

func Sqrt(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.FromFloat64(math.Sqrt(f64(args[0])))
}
