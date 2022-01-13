package funcs

import (
	"math"

	"github.com/nkanaev/numb/pkg/dimension"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

func toRadian(val value.Value) value.Number {
	if value.Type(val) == value.TYPE_NUMBER {
		return val.(value.Number)
	} else if value.Type(val) == value.TYPE_UNIT {
		val := val.(value.Unit)
		if measure, _ := val.Units.Dimension().Measure(); measure != dimension.ANGLE {
			panic("expected angle unit")
		}
		return value.Number{Num: unit.Convert(val.Num, val.Units, unit.Must("rad"))}
	}
	panic("unsupported type: " + value.Type(val).String())
}

func asRadian(val value.Number) value.Value {
	return value.Unit{Num: val.Num, Units: unit.Must("rad")}
}

func Sin(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Sin(f64(arg))), nil
}

func Cos(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Cos(f64(arg))), nil
}

func Tan(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Cos(f64(arg))), nil
}

func Asin(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	return asRadian(value.Float64(math.Asin(f64(args[0])))), nil
}

func Acos(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	return asRadian(value.Float64(math.Acos(f64(args[0])))), nil
}

func Atan(args ...value.Value) (value.Value, error) {
	arity(1, len(args))
	return asRadian(value.Float64(math.Atan(f64(args[0])))), nil
}
