package funcs

import (
	"fmt"
	"math"

	"github.com/nkanaev/numb/pkg/value"
)

func logF(fn func(float64) float64, args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}
	arg := args[0]

	num, ok := arg.(value.Number)
	if !ok {
		return nil, fmt.Errorf("expected number, got %s", value.Type(arg))
	}

	f64, exact := num.Num.Float64()
	if !exact && math.IsInf(f64, 0) {
		return nil, fmt.Errorf("%s magnitude is too large", arg)
	}

	return value.Float64(fn(f64)), nil
}

func Log(args ...value.Value) (value.Value, error) {
	return logF(math.Log, args...)
}

func Log2(args ...value.Value) (value.Value, error) {
	return logF(math.Log2, args...)
}

func Log10(args ...value.Value) (value.Value, error) {
	return logF(math.Log10, args...)
}

func Sqrt(args ...value.Value) (value.Value, error) {
	return logF(math.Sqrt, args...)
}
