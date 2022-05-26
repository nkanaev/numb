package funcs

import (
	"fmt"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/value"
)

func GCD(args ...value.Value) (value.Value, error) {
	if len(args) == 0 {
		return value.Int64(0), nil
	}
	var ret *big.Int
	for _, arg := range args {
		if value.Type(arg) != value.TYPE_NUMBER {
			return nil, fmt.Errorf("can accept only numbers: %s", arg)
		}
		num := arg.(value.Number).Num
		if !num.IsInt() {
			return nil, fmt.Errorf("not integer: %s", arg)
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
	return value.Number{Num: num}, nil
}

func LCM(args ...value.Value) (value.Value, error) {
	if len(args) == 0 {
		return value.Int64(0), nil
	}
	var ret *big.Int
	for _, arg := range args {
		if value.Type(arg) != value.TYPE_NUMBER {
			return nil, fmt.Errorf("can accept only numbers: %s", arg)
		}
		num := arg.(value.Number).Num
		if !num.IsInt() {
			return nil, fmt.Errorf("not integer: %s", arg)
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
	return value.Number{Num: num}, nil
}

func Abs(args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		return nil, fmt.Errorf("%s is not a number type", arg)
	}
	num := new(big.Rat).Set(arg.(value.Number).Num)
	if num.Cmp(ratutils.ZERO) == -1 {
		num.Neg(num)
	}
	return value.Number{Num: num}, nil
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

func Ceil(args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		return nil, fmt.Errorf("%s is not a number type", arg.String())
	}
	num := arg.(value.Number).Num
	return value.Number{Num: ceil(num)}, nil
}

func Floor(args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected one argument")
	}
	arg := args[0]
	if value.Type(arg) != value.TYPE_NUMBER {
		return nil, fmt.Errorf("%s is not a number type", arg)
	}
	num := arg.(value.Number).Num
	if num.IsInt() {
		return value.Number{Num: num}, nil
	}
	num = ceil(num)
	num.Sub(num, ratutils.ONE)
	return value.Number{Num: num}, nil
}

func Trunc(args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected one argument")
	}
	arg := args[0]
	switch arg.(type) {
	case value.Number:
		arg := arg.(value.Number)
		return value.Number{Num: ratutils.Trunc(arg.Num)}, nil
	case value.Unit:
		arg := arg.(value.Unit)
		ret := value.Unit{
			Num:   ratutils.Trunc(arg.Num),
			Units: arg.Units,
			Fmt:   arg.Fmt,
		}
		return ret, nil
	}
	return nil, fmt.Errorf("%s: unsupported type (%s)", arg, value.Type(arg).String())
}

func Round(args ...value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected one argument")
	}
	arg := args[0]
	switch arg.(type) {
    case value.Number:
        arg := arg.(value.Number)
        sign := 1
        if arg.Num.Cmp(ratutils.ZERO) < 0 {
            sign = -1
        }
        tmp := value.Number{
            Num: new(big.Rat).Add(arg.Num, big.NewRat(1*int64(sign), 2)),
            Fmt: arg.Fmt,
        }
        return Trunc(tmp)
	case value.Unit:
		arg := arg.(value.Unit)
        sign := 1
        if arg.Num.Cmp(ratutils.ZERO) < 0 {
            sign = -1
        }
		tmp := value.Unit{
            Num: new(big.Rat).Add(arg.Num, big.NewRat(1*int64(sign), 2)),
            Units: arg.Units,
            Fmt: arg.Fmt,
        }
        return Trunc(tmp)
	}
	return nil, fmt.Errorf("%s: unsupported type (%s)", arg, value.Type(arg).String())
}
