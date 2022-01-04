package value

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type ConformanceError struct {
	a, b unit.UnitList
}

func (c ConformanceError) Error() string {
	dim1, _ := c.a.Dimension().Measure()
	dim2, _ := c.b.Dimension().Measure()

	return fmt.Sprintf(
		"%s (%s) does not conform %s (%s)",
		c.a.String(), dim1.String(),
		c.b.String(), dim2.String())
}

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
			if !a.Units.Conforms(b.Units) {
				return nil, ConformanceError{a.Units, b.Units}
			}
			bnum := unit.Convert(b.Num, b.Units, a.Units)
			return Unit{Num: new(big.Rat).Add(a.Num, bnum), Units: a.Units}, nil
		case token.SUB:
			if !a.Units.Conforms(b.Units) {
				return nil, ConformanceError{a.Units, b.Units}
			}
			bnum := unit.Convert(b.Num, b.Units, a.Units)
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
			return Unit{Num: unit.Convert(a.Num, a.Units, b.Units), Units: b.Units}, nil
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
