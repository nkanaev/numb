package funcs

import (
	"math"
	"math/big"

	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

type mathOp func(float64) float64

func trigOp1(name string, op mathOp, args ...value.Value) value.Value {
	if len(args) != 1 {
		panic(name + ": expected 1 argument")
	}
	arg := args[0]
	if arg.Unit == nil {
		panic(name + ": provide rad or deg unit")
	}
	u := arg.Unit.String()
	if u != "deg" && u != "rad" {
		panic(name + ": expected rad or deg unit")
	}
	if u == "deg" {
		arg = arg.To(unit.Get("rad"))
	}
	f, exact := arg.Num.Float64()
	if !exact && math.IsInf(f, 0) {
		panic(name + ": value too large")
	}
	num := new(big.Rat).SetFloat64(op(f))
	return value.Value{Num: num, Unit: arg.Unit}
}

func Sin(args ...value.Value) value.Value {
	return trigOp1("sin", math.Sin, args...)
}

func Cos(args ...value.Value) value.Value {
	return trigOp1("cos", math.Cos, args...)
}

func Tan(args ...value.Value) value.Value {
	return trigOp1("tan", math.Tan, args...)
}

func Asin(args ...value.Value) value.Value {
	return trigOp1("asin", math.Asin, args...)
}

func Acos(args ...value.Value) value.Value {
	return trigOp1("acos", math.Acos, args...)
}

func Atan(args ...value.Value) value.Value {
	return trigOp1("atan", math.Atan, args...)
}
