package value

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type ValueType int

const (
	TYPE_UNKNOWN ValueType = iota
	TYPE_NUMBER
	TYPE_UNIT
	TYPE_NAME
)

type NumberFormat int

const (
	DEC NumberFormat = iota
	HEX
	OCT
	BIN
	RAT
	SCI
)

func (t ValueType) String() string {
	switch t {
	case TYPE_NUMBER:
		return "number"
	case TYPE_UNIT:
		return "unit"
	case TYPE_NAME:
		return "name"
	default:
		return "unknown"
	}
}

func Type(x Value) ValueType {
	switch x.(type) {
	case Number:
		return TYPE_NUMBER
	case Unit:
		return TYPE_UNIT
	case Name:
		return TYPE_NAME
	default:
		return TYPE_UNKNOWN
	}
}

type Value interface {
	BinOP(token.Token, Value) (Value, error)
	UnOP(token.Token) (Value, error)
	String() string
}

type Number struct {
	Num *big.Rat
	Fmt NumberFormat
}

type IntOperationError struct {
	op token.Token
}

func (x IntOperationError) Error() string {
	return x.op.String() + " is only supported between integers"
}

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

type UnsupportedBinOP struct {
	a, b Value
	op token.Token
}

func (err UnsupportedBinOP) Error() string {
	return fmt.Sprintf(
		"unsupported operation `%s` between %s and %s",
		err.op.String(), Type(err.a).String(), Type(err.b).String(),
	)
}

func (a Number) String() string {
	return formatNum(a.Num, a.Fmt, "_", 2)
}

func (a Number) BinOP(op token.Token, b Value) (Value, error) {
	// TODO: int type casts lose precision
	switch b.(type) {
	case Number:
		b := b.(Number)
		var n *big.Rat
		switch op {
		case token.ADD:
			n = new(big.Rat).Add(a.Num, b.Num)
		case token.SUB:
			n = new(big.Rat).Add(a.Num, b.Num)
		case token.MUL:
			n = new(big.Rat).Add(a.Num, b.Num)
		case token.QUO:
			n = new(big.Rat).Add(a.Num, b.Num)
		case token.EXP:
			if !b.Num.IsInt() {
				return nil, errors.New("exponentiation does not support non-integer power")
			}
			n = ratutils.ExpInt(a.Num, int(b.Num.Num().Int64()))
		case token.REM:
			n = ratutils.Mod(a.Num, b.Num)
		case token.SHL:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			if b.Num.Cmp(ratutils.ZERO) < 0 {
				return nil, errors.New("negative shift")
			}
			n = new(big.Rat).Set(a.Num)
			n.Num().Lsh(n.Num(), uint(b.Num.Num().Uint64()))
		case token.SHR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			if b.Num.Cmp(ratutils.ZERO) < 0 {
				return nil, errors.New("negative shift")
			}
			n = new(big.Rat).Set(a.Num)
			n.Num().Rsh(n.Num(), uint(b.Num.Num().Uint64()))
		case token.AND:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			n = new(big.Rat).Set(a.Num)
			n.Num().And(n.Num(), b.Num.Num())
		case token.OR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			n = new(big.Rat).Set(a.Num)
			n.Num().Or(n.Num(), b.Num.Num())
		case token.XOR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			n := new(big.Rat).Set(a.Num)
			n.Num().Xor(n.Num(), b.Num.Num())
		}
		if n != nil {
			return Number{Num: n, Fmt: a.Fmt}, nil
		}
	case Unit:
		b := b.(Unit)
		switch op {
		case token.MUL:
			return Unit{Num: new(big.Rat).Mul(a.Num, b.Num), Units: b.Units}, nil
		case token.QUO:
			return Unit{Num: new(big.Rat).Quo(a.Num, b.Num), Units: b.Units.Exp(-1)}, nil
		}
	}
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (a Number) UnOP(op token.Token) (Value, error) {
	if op == token.SUB {
		return Number{Num: new(big.Rat).Neg(a.Num)}, nil
	}
	return nil, errors.New("unsupported unary operation: %s" + op.String())
}

type Unit struct {
	Num   *big.Rat
	Units unit.UnitList
}

func (a Unit) String() string {
	return formatNum(a.Num, DEC, "_", 2) + " " + a.Units.String()
}

func (a Unit) BinOP(op token.Token, b Value) (Value, error) {
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

type Name struct {
	Val string
}

func (n Name) BinOP(t token.Token, v Value) (Value, error) {
	return nil, errors.New("name cannot be used for operations")
}

func (n Name) UnOP(t token.Token) (Value, error) {
	return nil, errors.New("name cannot be used for operations")
}

func (n Name) String() string {
	return n.Val
}

func Int64(x int64) Value {
	return Number{Num: new(big.Rat).SetInt64(x)}
}

func Float64(x float64) Value {
	return Number{Num: new(big.Rat).SetFloat64(x)}
}
