package funcs

import (
	"math"
	"math/big"

	"github.com/nkanaev/numb/pkg/consts"
	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/unit/dimension"
	"github.com/nkanaev/numb/pkg/value"
)

type mathOp func(float64) float64

var radian = unit.Must("rad")

func trigOp1(name string, op mathOp, args ...value.Value) value.Value {
	if len(args) != 1 {
		panic(name + ": expected 1 argument")
	}
	arg := args[0]
	if len(arg.Unit) == 0 || !arg.Unit.Dimension().Equals(dimension.ANGLE.Dim()) {
		panic(name + ": can accept only dimensions of angle")
	}
	arg = arg.To(radian)
	// TODO: check for negative values
	x := ratutils.ModRat(arg.Num, consts.PI)
	f, exact := x.Float64()
	if !exact && math.IsInf(f, 0) {
		panic(name + ": value too large")
	}
	num := new(big.Rat).SetFloat64(op(f))
	return value.Value{Num: num}
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
