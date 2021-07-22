package funcs

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/value"
)

func GCD(args ...value.Value) value.Value {
	if len(args) == 0 {
		return value.NewInt(0)
	}
	var ret *big.Int
	for _, arg := range args {
		if len(arg.Unit) > 0 {
			panic("gcd: cannot accept argument with unit: " + arg.String())
		}
		argint, isint := ratutils.ToInt(arg.Num)
		if !isint {
			panic("gcd: not integer: " + arg.String())
		}
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
