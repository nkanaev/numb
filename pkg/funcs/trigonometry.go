package funcs

import (
	"math"

	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/dimension"
	"github.com/nkanaev/numb/pkg/value"
)

func toRadian(val value.Value) value.Value {
	if value.Type(val) == value.TYPE_NUMBER {
		return value.Unit{Num: val.(value.Number).Num, Units: unit.Must("rad")}
	}
	if value.Type(val) == value.TYPE_UNIT {
		val := val.(value.Unit)
		if measure, _ := val.Units.Dimension().Measure(); measure != dimension.ANGLE {
			panic("expected angle unit")
		}
		return value.Unit{Num: unit.Convert(val.Num, val.Units, unit.Must("rad"))}
	}
	panic("unsupported type: " + value.Type(val).String())
}

func Sin(args ...value.Value) value.Value {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Sin(f64(arg)))
}

func Cos(args ...value.Value) value.Value {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Cos(f64(arg)))
}

func Tan(args ...value.Value) value.Value {
	arity(1, len(args))
	arg := toRadian(args[0])
	return value.Float64(math.Cos(f64(arg)))
}

func Asin(args ...value.Value) value.Value {
	arity(1, len(args))
	return toRadian(value.Float64(math.Asin(f64(args[0]))))
}

func Acos(args ...value.Value) value.Value {
	arity(1, len(args))
	return toRadian(value.Float64(math.Acos(f64(args[0]))))
}

func Atan(args ...value.Value) value.Value {
	arity(1, len(args))
	return toRadian(value.Float64(math.Atan(f64(args[0]))))
}
