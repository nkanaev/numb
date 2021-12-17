package funcs

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/value"
)

func GCD(args ...value.Value) value.Value {
	if len(args) == 0 {
		return value.Int64(0)
	}
	var ret *big.Int
	for _, arg := range args {
		if value.Type(arg) != value.TYPE_NUMBER {
			panic("can accept only numbers: " + arg.String())
		}
		num := arg.(value.Number).Num
		if !num.IsInt() {
			panic("not integer: " + arg.String())
		}
		argint := ratutils.TruncInt(num)
		if ret == nil {
			ret = argint
		} else {
			ret.GCD(nil, nil, ret, argint)
		}
	}
	num := big.NewRat(1, 1)
	num.Num().Set(ret)
	return value.Number{Num: num}
}

func LCM(args ...value.Value) value.Value {
	if len(args) == 0 {
		return value.Int64(0)
	}
	var ret *big.Int
	for _, arg := range args {
		if value.Type(arg) != value.TYPE_NUMBER {
			panic("can accept only numbers: " + arg.String())
		}
		num := arg.(value.Number).Num
		if !num.IsInt() {
			panic("not integer: " + arg.String())
		}
		argint := ratutils.TruncInt(num)
		if ret == nil {
			ret = argint
		} else {
			gcd := new(big.Int).Set(ret)
			gcd.GCD(nil, nil, gcd, argint)

			ret.Mul(ret, argint)
			ret.Quo(ret, gcd)
		}
	}
	num := big.NewRat(1, 1)
	num.Num().Set(ret)
	return value.Number{Num: num}
}

func Abs(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("abs: expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		panic("can accept only numbers: " + arg.String())
	}
	num := new(big.Rat).Set(arg.(value.Number).Num)
	if num.Cmp(ratutils.ZERO) == -1 {
		num.Neg(num)
	}
	return value.Number{Num: num}
}

func ceil(num *big.Rat) *big.Rat {
	num = new(big.Rat).Set(num)
	if num.IsInt() {
		return num
	}
	if num.Cmp(ratutils.ZERO) > 0 {
		num = ratutils.Trunc(num)
		num.Add(num, ratutils.ONE)
		return num
	}
	num = ratutils.Trunc(num)
	return num
}

func Ceil(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("ceil: expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		panic("can accept only numbers: " + arg.String())
	}
	num := arg.(value.Number).Num
	return value.Number{Num: ceil(num)}
}

func Floor(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("floor: expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		panic("can accept only numbers: " + arg.String())
	}
	num := arg.(value.Number).Num
	if num.IsInt() {
		return value.Number{Num: num}
	}
	num = ceil(num)
	num.Sub(num, ratutils.ONE)
	return value.Number{Num: num}
}

func Trunc(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("trunc: expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		panic("can accept only numbers: " + arg.String())
	}
	num := arg.(value.Number).Num
	return value.Number{Num: ratutils.Trunc(num)}
}
