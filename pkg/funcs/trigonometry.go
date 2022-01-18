package funcs

import (
	"fmt"
	"math"
	"math/big"

	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

func trigOp(fn func(float64) float64, inUnit, outUnit bool, args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}
	arg := args[0]

	var num *big.Rat
	switch arg.(type) {
	case value.Number:
		num = arg.(value.Number).Num
	case value.Unit:
		if !inUnit {
			return nil, fmt.Errorf("expected number, got %s", value.Type(arg))
		}
		arg := arg.(value.Unit)
		var err error
		num, err = unit.Convert(arg.Num, arg.Units, unit.Must("rad"))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported type: %s", value.Type(arg))
	}

	f64, exact := num.Float64()
	if !exact && math.IsInf(f64, 0) {
		return nil, fmt.Errorf("%s magnitude is too large", arg)
	}

	if outUnit {
		return value.Unit{
			Num: new(big.Rat).SetFloat64(fn(f64)),
			Units: unit.Must("rad"),
		}, nil
	}
	return value.Float64(fn(f64)), nil
}

func Sin(args ...value.Value) (value.Value, error) {
	return trigOp(math.Sin, true, false, args...)
}

func Cos(args ...value.Value) (value.Value, error) {
	return trigOp(math.Cos, true, false, args...)
}

func Tan(args ...value.Value) (value.Value, error) {
	return trigOp(math.Tan, true, false, args...)
}

func Asin(args ...value.Value) (value.Value, error) {
	return trigOp(math.Asin, false, true, args...)
}

func Acos(args ...value.Value) (value.Value, error) {
	return trigOp(math.Acos, false, true, args...)
}

func Atan(args ...value.Value) (value.Value, error) {
	return trigOp(math.Atan, false, true, args...)
}
