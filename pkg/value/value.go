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

// TODO: fix all Rat.Num() to int(Rat)

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

func (a Value) Mul(b Value) Value {
	return Value{Num: new(big.Rat).Mul(a.Num, b.Num), Fmt: a.Fmt}
}

func (a Value) Add(b Value) Value {
	return Value{Num: new(big.Rat).Add(a.Num, b.Num), Fmt: a.Fmt}
}

func (a Value) Sub(b Value) Value {
	return Value{Num: new(big.Rat).Sub(a.Num, b.Num), Fmt: a.Fmt}
}

func (a Value) Quo(b Value) Value {
	return Value{Num: new(big.Rat).Quo(a.Num, b.Num), Fmt: a.Fmt}
}

func (a Value) Lsh(b Value) Value {
	ia, ib := toInt(a.Num), uint(toInt(b.Num).Uint64())
	num := big.NewRat(1, 1)
	num.Num().Lsh(ia, ib)
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Rsh(b Value) Value {
	ia, ib := toInt(a.Num), uint(toInt(b.Num).Uint64())
	num := big.NewRat(1, 1)
	num.Num().Rsh(ia, ib)
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) And(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().And(toInt(a.Num), toInt(b.Num))
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Or(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Or(toInt(a.Num), toInt(b.Num))
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Xor(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Xor(toInt(a.Num), toInt(b.Num))
	return Value{Num: num, Fmt: a.Fmt}
}

func (a Value) Rem(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Rem(toInt(a.Num), toInt(b.Num))
	return Value{Num: num, Fmt: a.Fmt}
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
