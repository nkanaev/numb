package value

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/unit"
)

type Value struct {
	Num  *big.Rat
	Fmt  NumeralSystem
	Unit unit.UnitList
}

func Int64(x int64) Value {
	return Value{Num: big.NewRat(x, 1)}
}

func Float64(x float64) Value {
	return Value{Num: new(big.Rat).SetFloat64(x)}
}

func Parse(x string) Value {
	num, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse number: " + x)
	}
	return Value{Num: num}
}

func prepare(a, b Value) (Value, Value, unit.UnitList) {
	var u unit.UnitList
	if a.Unit != nil {
		u = a.Unit
		if len(b.Unit) > 0 {
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
	int := op(ratutils.ToInt(a.Num), ratutils.ToInt(b.Num))
	rat := big.NewRat(1, 1)
	rat.Num().Set(int)
	return Value{Num: rat, Fmt: a.Fmt, Unit: u}
}

func doShift(a, b Value, op func(*big.Int, uint) *big.Int) Value {
	a, b, u := prepare(a, b)
	ia, ib := ratutils.ToInt(a.Num), uint(ratutils.ToInt(b.Num).Uint64())
	num := big.NewRat(1, 1)
	num.Num().Set(op(ia, ib))
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
}

func (a Value) Add(b Value) Value {
	return do(a, b, new(big.Rat).Add)
}

func (a Value) Sub(b Value) Value {
	return do(a, b, new(big.Rat).Sub)
}

func (a Value) Mul(b Value) Value {
	n := new(big.Rat).Mul(a.Num, b.Num)

	u := a.Unit
	if u == nil {
		u = b.Unit
	}
	if a.Unit != nil && b.Unit != nil {
		utmp := a.Unit.Mul(b.Unit)
		u = utmp.Simplify()
		n = unit.Convert(n, utmp, u)
	}

	return Value{Num: n, Fmt: a.Fmt, Unit: u}
}

func (a Value) Quo(b Value) Value {
	n := new(big.Rat).Quo(a.Num, b.Num)

	u := a.Unit
	if u == nil {
		u = b.Unit
	}
	if a.Unit != nil && b.Unit != nil {
		utmp := a.Unit.Quo(b.Unit)
		u = utmp.Simplify()
		n = unit.Convert(n, utmp, u)
	}

	return Value{Num: n, Fmt: a.Fmt, Unit: u}
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
	n := ratutils.Mod(a.Num, b.Num)

	u := a.Unit
	if u == nil {
		u = b.Unit
	}
	if a.Unit != nil && b.Unit != nil {
		utmp := a.Unit.Quo(b.Unit)
		u = utmp.Simplify()
		n = unit.Convert(n, utmp, u)
	}

	return Value{Num: n, Fmt: a.Fmt, Unit: u}
}

func (a Value) Exp(b Value) Value {
	a, b, u := prepare(a, b)
	exp := int(ratutils.ToInt(b.Num).Int64())
	num := ratutils.ExpInt(a.Num, exp)
	if u != nil {
		u = u.Exp(exp)
	}
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
}

func (a Value) Neg() Value {
	return Value{Num: new(big.Rat).Neg(a.Num), Fmt: a.Fmt, Unit: a.Unit}
}

func (a Value) As(n NumeralSystem) Value {
	a.Fmt = n
	return a
}

func (a Value) To(u unit.UnitList) Value {
	num := unit.Convert(a.Num, a.Unit, u)
	return Value{Num: num, Fmt: a.Fmt, Unit: u}
}

func (a Value) WithUnit(u unit.UnitList) Value {
	a.Unit = u
	return a
}

func (a Value) WithFormat(fmt NumeralSystem) Value {
	a.Fmt = fmt
	return a
}

func (a Value) Eval(map[string]Value) Value {
	return a
}
