package value

import (
	"errors"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type Unit struct {
	Num   *big.Rat
	Fmt   NumberFormat
	Units unit.Units
}

func (a Unit) String() string {
	return formatNum(a.Num, a.Fmt, defaultSep, defaultPrec) + " " + a.Units.String()
}

func (a Unit) BinOP(op token.Token, b Value) (Value, error) {
	switch b.(type) {
	case Number:
		bnum := b.(Number).Num
		switch op {
		case token.MUL:
			num := new(big.Rat).Mul(a.Num, bnum)
			return Unit{Num: num, Units: a.Units, Fmt: a.Fmt}, nil
		case token.QUO:
			num := new(big.Rat).Quo(a.Num, bnum)
			return Unit{Num: num, Units: a.Units, Fmt: a.Fmt}, nil
		case token.EXP:
			if !bnum.IsInt() {
				return nil, errors.New("exponentiation does not support non-integer power")
			}
			exp := int(bnum.Num().Int64())
			n := ratutils.ExpInt(a.Num, exp)
			u := a.Units.Exp(exp)
			if u.Dimension().IsPure() {
				return Number{Num: n}, nil
			}
			return Unit{Num: n, Units: u, Fmt: a.Fmt}, nil
		}
	case Unit:
		b := b.(Unit)
		switch op {
		case token.ADD:
			bnum, err := unit.Convert(b.Num, b.Units, a.Units)
			if err != nil {
				return nil, err
			}
			num := new(big.Rat).Add(a.Num, bnum)
			return Unit{Num: num, Units: a.Units, Fmt: a.Fmt}, nil
		case token.SUB:
			bnum, err := unit.Convert(b.Num, b.Units, a.Units)
			if err != nil {
				return nil, err
			}
			num := new(big.Rat).Sub(a.Num, bnum)
			return Unit{Num: num, Units: a.Units, Fmt: a.Fmt}, nil
		case token.MUL:
			tmpu := a.Units.Mul(b.Units)
			newu := tmpu.Simplify()

			tmpn := new(big.Rat).Mul(a.Num, b.Num)
			newn, _ := unit.Convert(tmpn, tmpu, newu)

			if newu.Dimension().IsPure() {
				return Number{Num: newn}, nil
			}
			return Unit{Num: newn, Units: newu, Fmt: a.Fmt}, nil
		case token.QUO:
			tmpu := a.Units.Quo(b.Units)
			newu := tmpu.Simplify()

			tmpn := new(big.Rat).Quo(a.Num, b.Num)
			newn, _ := unit.Convert(tmpn, tmpu, newu)

			if newu.Dimension().IsPure() {
				return Number{Num: newn}, nil
			}
			return Unit{Num: newn, Units: newu, Fmt: a.Fmt}, nil
		case token.TO:
			if b.Num.Cmp(ratutils.ONE) != 0 {
				return nil, errors.New("cannot convert to a unit with a value: " + b.String())
			}
			num, err := unit.Convert(a.Num, a.Units, b.Units)
			if err != nil {
				return nil, err
			}
			return Unit{Num: num, Units: b.Units, Fmt: a.Fmt}, nil
		}
	}
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (a Unit) UnOP(op token.Token) (Value, error) {
	if op == token.SUB {
		num := new(big.Rat).Neg(a.Num)
		return Unit{Num: num, Units: a.Units, Fmt: a.Fmt}, nil
	}
	return nil, errors.New("unsupported unary operation: %s" + op.String())
}

func (a Unit) In(fmt string) (Value, error) {
	f, ok := StringToFormat[fmt]
	if !ok {
		return nil, errors.New("unrecognized number format for unit: " + fmt)
	}
	return Unit{Num: a.Num, Fmt: f, Units: a.Units}, nil
}
