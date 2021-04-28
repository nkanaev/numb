package value

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/consts"
	"github.com/nkanaev/numb/pkg/unit"
)

type Value struct {
	Num  *big.Rat
	Fmt  NumeralSystem
	Prec int
	Unit *unit.Unit
}

var Consts = map[string]Value{
	"pi": Value{Num: consts.PI},
	"e":  Value{Num: consts.E},
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

func prepare(a, b Value) (Value, Value, *unit.Unit) {
	var u *unit.Unit
	if a.Unit != nil {
		u = a.Unit
		if b.Unit != nil {
			b = b.To(a.Unit)
		}
	} else if b.Unit != nil {
		u = b.Unit
	}
	return a, b, u
}

func do(a, b Value, op func(*big.Rat, *big.Rat) *big.Rat) Value {
	a, b, u := prepare(a, b)
	return Value{Num: op(a.Num, b.Num), Fmt: a.Fmt, Unit: u}
}

func doInt(a, b Value, op func(*big.Int, *big.Int) *big.Int) Value {
	a, b, u := prepare(a, b)
	int := op(toInt(a.Num), toInt(b.Num))
	rat := big.NewRat(1, 1)
	rat.Num().Set(int)
	return Value{Num: rat, Fmt: a.Fmt, Unit: u}
}

func doShift(a, b Value, op func(*big.Int, uint) *big.Int) Value {
	a, b, u := prepare(a, b)
	ia, ib := toInt(a.Num), uint(toInt(b.Num).Uint64())
	num := big.NewRat(1, 1)
	num.Num().Set(op(ia, ib))
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
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
	a, b, u := prepare(a, b)
	num := big.NewRat(1, 1)
	num.Num().Exp(toInt(a.Num), toInt(b.Num), nil)
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
}

func (a Value) Neg() Value {
	return Value{Num: new(big.Rat).Neg(a.Num), Fmt: a.Fmt, Unit: a.Unit}
}

func (a Value) As(n NumeralSystem) Value {
	a.Fmt = n
	return a
}

func (a Value) To(u *unit.Unit) Value {
	num := unit.Convert(a.Num, a.Unit, u)
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
}

func (a Value) WithUnit(u *unit.Unit) Value {
	a.Unit = u
	return a
}

func (a Value) WithPrec(p int) Value {
	a.Prec = p
	return a
}

func (a Value) String() string {
	num := ""
	switch a.Fmt {
	case DEC:
		if a.Num.IsInt() {
			num = a.Num.RatString()
		} else {
			prec := a.Prec
			if prec == 0 {
				prec = 4
			}
			// TODO: trailing zeros
			num = a.Num.FloatString(prec)
		}
	case HEX:
		num = fmt.Sprintf("%#x", toInt(a.Num))
	case OCT:
		num = fmt.Sprintf("%O", toInt(a.Num))
	case BIN:
		num = fmt.Sprintf("%#b", toInt(a.Num))
	case RAT:
		num = a.Num.String()
	case EXP:
		num = fmt.Sprintf("%e", new(big.Float).SetRat(a.Num))
	case WAT:
		suffixes := "KMGTPEZY"
		thousand := big.NewRat(1000, 1)

		if a.Num.Cmp(thousand) < 0 {
			return a.As(DEC).String()
		}

		x := new(big.Rat).Set(a.Num)
		var i int
		for ; i < len(suffixes) && x.Cmp(thousand) >= 0; i++ {
			x.Quo(x, thousand)
		}
		if x.IsInt() {
			num = x.RatString() + string(suffixes[i-1])
		} else {
			num = x.FloatString(1) + string(suffixes[i-1])
		}
	}
	if a.Unit != nil {
		num += " " + a.Unit.String()
	}
	return num
}

func (a Value) Eval(map[string]Value) Value {
	return a
}
