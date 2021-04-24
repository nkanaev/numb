package value

import (
	"fmt"
	"math/big"
	"strings"
)

type Value struct {
	Num  *big.Rat
	Fmt  NumeralSystem
	Prec int
}

func toInt(x *big.Rat) *big.Int {
	return new(big.Int).Div(x.Num(), x.Denom())
}

func NewInt(x int64) Value {
	return Value{Num: big.NewRat(x, 1)}
}

func Parse(x string) Value {
	num, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse number: " + x)
	}
	base := DEC
	if strings.HasPrefix(x, "0x") {
		base = HEX
	} else if strings.HasPrefix(x, "0o") {
		base = OCT
	} else if strings.HasPrefix(x, "0b") {
		base = BIN
	}
	return Value{Num: num, Fmt: base}
}

func do(a, b Value, op func(*big.Rat, *big.Rat) *big.Rat) Value {
	return Value{Num: op(a.Num, b.Num), Fmt: a.Fmt}
}

func doInt(a, b Value, op func(*big.Int, *big.Int) *big.Int) Value {
	int := op(toInt(a.Num), toInt(b.Num))
	rat := big.NewRat(1, 1)
	rat.Num().Set(int)
	return Value{Num: rat, Fmt: a.Fmt}
}

func doShift(a, b Value, op func(*big.Int, uint) *big.Int) Value {
	ia, ib := toInt(a.Num), uint(toInt(b.Num).Uint64())
	num := big.NewRat(1, 1)
	num.Num().Set(op(ia, ib))
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Mul(b Value) Value {
	return do(a, b, new(big.Rat).Mul)
}

func (a Value) Add(b Value) Value {
	return do(a, b, new(big.Rat).Add)
}

func (a Value) Sub(b Value) Value {
	return do(a, b, new(big.Rat).Sub)
}

func (a Value) Quo(b Value) Value {
	return do(a, b, new(big.Rat).Quo)
}

func (a Value) Lsh(b Value) Value {
	return doShift(a, b, new(big.Int).Lsh)
}

func (a Value) Rsh(b Value) Value {
	return doShift(a, b, new(big.Int).Rsh)
}

func (a Value) And(b Value) Value {
	return doInt(a, b, new(big.Int).And)
}

func (a Value) Or(b Value) Value {
	return doInt(a, b, new(big.Int).Or)
}

func (a Value) Xor(b Value) Value {
	return doInt(a, b, new(big.Int).Xor)
}

func (a Value) Rem(b Value) Value {
	return doInt(a, b, new(big.Int).Rem)
}

func (a Value) Exp(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Exp(toInt(a.Num), toInt(b.Num), nil)
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Neg() Value {
	return Value{Num: new(big.Rat).Neg(a.Num), Fmt: a.Fmt}
}

func (a Value) As(n NumeralSystem) Value {
	a.Fmt = n
	return a
}

func (a Value) WithPrec(p int) Value {
	a.Prec = p
	return a
}

func (a Value) String() string {
	switch a.Fmt {
	case DEC:
		if a.Num.IsInt() {
			return a.Num.RatString()
		}
		prec := a.Prec
		if prec == 0 {
			prec = 4
		}
		// TODO: trailing zeros
		return a.Num.FloatString(prec)
	case HEX:
		return fmt.Sprintf("%#x", toInt(a.Num))
	case OCT:
		return fmt.Sprintf("%O", toInt(a.Num))
	case BIN:
		return fmt.Sprintf("%#b", toInt(a.Num))
	case RAT:
		return a.Num.String()
	}
	return a.Num.String()
}

func (a Value) Eval(map[string]Value) Value {
	return a
}
