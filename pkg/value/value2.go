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

func Type(x Value2) ValueType {
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

type Value2 interface {
	BinOP(token.Token, Value2) (Value2, error)
	UnOP(token.Token) (Value2, error)
	String() string
}

type Number struct {
	Num *big.Rat
	Fmt string
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

func (a Number) String() string {
	return a.Num.String()
}

func (a Number) BinOP(op token.Token, b Value2) (Value2, error) {
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
			// TODO: exponent
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
	return nil, errors.New("unsupported operation: " + op.String())
}

func (a Number) UnOP(op token.Token) (Value2, error) {
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
	return a.Num.String() + " " + a.Units.String()
}

func (a Unit) BinOP(op token.Token, b Value2) (Value2, error) {
	switch b.(type) {
	case Number:
		bnum := b.(Number).Num
		switch op {
		case token.MUL:
			return Unit{Num: new(big.Rat).Mul(a.Num, bnum), Units: a.Units}, nil
		case token.QUO:
			return Unit{Num: new(big.Rat).Quo(a.Num, bnum), Units: a.Units}, nil
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
	return nil, errors.New("Unsupported operation")
}

func (a Unit) UnOP(op token.Token) (Value2, error) {
	if op == token.SUB {
		return Unit{Num: new(big.Rat).Neg(a.Num), Units: a.Units}, nil
	}
	return nil, errors.New("unsupported unary operation: %s" + op.String())
}

type Name struct {
	Val string
}

func (n Name) BinOP(t token.Token, v Value2) (Value2, error) {
	return nil, errors.New("name cannot be used for operations")
}

func (n Name) UnOP(t token.Token) (Value2, error) {
	return nil, errors.New("name cannot be used for operations")
}

func (n Name) String() string {
	return n.Val
}
