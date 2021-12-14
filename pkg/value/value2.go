package value

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type Value2 interface {
	Do(token.Token, Value2) (Value2, error)
	String() string
}

type Number struct {
	Num *big.Rat
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

func (a Number) Do(op token.Token, b Value2) (Value2, error) {
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
			num := new(big.Rat).Set(a.Num)
			num.Num().Lsh(num.Num(), uint(b.Num.Num().Uint64()))
			return Number{Num: num}, nil
		case token.SHR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			if b.Num.Cmp(ratutils.ZERO) < 0 {
				return nil, errors.New("negative shift")
			}
			num := new(big.Rat).Set(a.Num)
			num.Num().Rsh(num.Num(), uint(b.Num.Num().Uint64()))
			return Number{Num: num}, nil
		case token.AND:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			num := new(big.Rat).Set(a.Num)
			num.Num().And(num.Num(), b.Num.Num())
			return Number{Num: num}, nil
		case token.OR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			num := new(big.Rat).Set(a.Num)
			num.Num().Or(num.Num(), b.Num.Num())
			return Number{Num: num}, nil
		case token.XOR:
			if !(a.Num.IsInt() && b.Num.IsInt()) {
				return nil, IntOperationError{op}
			}
			num := new(big.Rat).Set(a.Num)
			num.Num().Xor(num.Num(), b.Num.Num())
			return Number{Num: num}, nil
		}
		if n != nil {
			return Number{Num: n}, nil
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

type Unit struct {
	Num   *big.Rat
	Units unit.UnitList
}

func (a Unit) String() string {
	return a.Num.String() + " " + a.Units.String()
}

func (a Unit) Do(op token.Token, b Value2) (Value2, error) {
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
		}
	}
	return nil, errors.New("Unsupported operation")
}
