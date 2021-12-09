package funcs

import (
	"math"

	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/dimension"
	"github.com/nkanaev/numb/pkg/value"
)

func toRadian(val value.Value) value.Value {
	if val.Unit.Dimension().IsZero() {
		return val.WithUnit(unit.Must("rad"))
	}
	if measure, _ := val.Unit.Dimension().Measure(); measure != dimension.ANGLE {
		panic("expected angle unit")
	}
	return val.To(unit.Must("rad"))
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
	return value.Float64(math.Asin(f64(args[0]))).WithUnit(unit.Must("rad"))
}

func Acos(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.Float64(math.Acos(f64(args[0]))).WithUnit(unit.Must("rad"))
}

func Atan(args ...value.Value) value.Value {
	arity(1, len(args))
	return value.Float64(math.Atan(f64(args[0]))).WithUnit(unit.Must("rad"))
}
