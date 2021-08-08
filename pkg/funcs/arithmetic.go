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
		if len(arg.Unit) > 0 {
			panic("cannot accept argument with unit: " + arg.String())
		}
		if !ratutils.IsInt(arg.Num) {
			panic("not integer: " + arg.String())
		}
		argint := ratutils.ToInt(arg.Num)
		if ret == nil {
			ret = argint
		} else {
			ret.GCD(nil, nil, ret, argint)
		}
	}
	num := big.NewRat(1, 1)
	num.Num().Set(ret)
	return value.Value{Num: num}
}

func LCM(args ...value.Value) value.Value {
	if len(args) == 0 {
		return value.Int64(0)
	}
	var ret *big.Int
	for _, arg := range args {
		if len(arg.Unit) > 0 {
			panic("lcm: cannot accept argument with unit: " + arg.String())
		}
		if !ratutils.IsInt(arg.Num) {
			panic("lcm: not integer: " + arg.String())
		}
		argint := ratutils.ToInt(arg.Num)
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
	return value.Value{Num: num}
}

func Abs(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("abs: expected one argument")
	}
	num := new(big.Rat).Set(args[0].Num)
	if num.Cmp(ratutils.ZERO) == -1 {
		num.Neg(num)
	}
	return value.Value{Num: num}
}

func Ceil(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("ceil: expected one argument")
	}
	arg := args[0]
	if arg.Num.IsInt() {
		return value.Value{Num: arg.Num}
	}
	int := ratutils.ToInt(arg.Num)
	if arg.Num.Cmp(ratutils.ZERO) > 0 {
		int.Add(int, ratutils.ONE.Num())
	}
	return value.Value{Num: ratutils.FromInt(int)}
}

func Floor(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("floor: expected one argument")
	}
	arg := args[0]
	if arg.Num.IsInt() {
		return value.Value{Num: arg.Num}
	}
	ret := Ceil(arg)
	ret.Num.Sub(ret.Num, ratutils.ONE)
	return ret
}

func Trunc(args ...value.Value) value.Value {
	if len(args) != 1 {
		panic("trunc: expected one argument")
	}
	return value.Value{Num: ratutils.Trunc(args[0].Num)}
}
