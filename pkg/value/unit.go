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
	Units unit.UnitList
}

func (a Unit) String() string {
	return formatNum(a.Num, a.Fmt, ",", 2) + " " + a.Units.String()
}

func (a Unit) BinOP(op token.Token, b Value) (Value, error) {
	// TODO: preserve format
	switch b.(type) {
	case Number:
		bnum := b.(Number).Num
		switch op {
		case token.MUL:
			return Unit{Num: new(big.Rat).Mul(a.Num, bnum), Units: a.Units}, nil
		case token.QUO:
			return Unit{Num: new(big.Rat).Quo(a.Num, bnum), Units: a.Units}, nil
		case token.EXP:
			if !bnum.IsInt() {
				return nil, errors.New("exponentiation does not support non-integer power")
			}
			exp := int(bnum.Num().Int64())
			n := ratutils.ExpInt(a.Num, exp)
			u := a.Units.Exp(exp)
			if u.Dimension().IsZero() {
				return Number{Num: n}, nil
			}
			return Unit{Num: n, Units: u}, nil
		}
	case Unit:
		b := b.(Unit)
		switch op {
		case token.ADD:
			bnum, err := unit.Convert(b.Num, b.Units, a.Units)
			if err != nil {
				return nil, err
			}
			return Unit{Num: new(big.Rat).Add(a.Num, bnum), Units: a.Units}, nil
		case token.SUB:
			bnum, err := unit.Convert(b.Num, b.Units, a.Units)
			if err != nil {
				return nil, err
			}
			return Unit{Num: new(big.Rat).Sub(a.Num, bnum), Units: a.Units}, nil
		case token.MUL:
			newn := new(big.Rat).Mul(a.Num, b.Num)
			newu := a.Units.Mul(b.Units)
			if newu.Dimension().IsZero() {
				return Number{Num: newn}, nil
			}
			return Unit{Num: newn, Units: newu}, nil
		case token.QUO:
			newn := new(big.Rat).Quo(a.Num, b.Num)
			newu := a.Units.Quo(b.Units)
			if newu.Dimension().IsZero() {
				return Number{Num: newn}, nil
			}
			return Unit{Num: newn, Units: newu}, nil
		case token.TO:
			if b.Num.Cmp(ratutils.ONE) != 0 {
				return nil, errors.New("cannot convert to a unit with a value: " + b.String())
			}
			num, err := unit.Convert(a.Num, a.Units, b.Units)
			if err != nil {
				return nil, err
			}
			return Unit{Num: num, Units: b.Units}, nil
		}
	case Percent:
		switch op {
		case token.ADD, token.SUB, token.MUL, token.QUO:
			b := b.(Percent)
			bnum := b.Apply(a.Num)
			ret, err := Number{Num: a.Num}.BinOP(op, Number{Num: bnum})
			if err != nil {
				if errors.Is(err, UnsupportedBinOP{}) {
					return nil, UnsupportedBinOP{a: a, b: b, op: op}
				}
				return nil, err
			}
			return Unit{Num: ret.(Number).Num, Units: a.Units, Fmt: a.Fmt}, err
		}
	}
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (a Unit) UnOP(op token.Token) (Value, error) {
	if op == token.SUB {
		return Unit{Num: new(big.Rat).Neg(a.Num), Units: a.Units}, nil
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
