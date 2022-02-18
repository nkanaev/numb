package value

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type IntOperationError struct {
	op token.Token
}

func (x IntOperationError) Error() string {
	return x.op.String() + " is only supported between integers"
}

type UnsupportedBinOP struct {
	a, b Value
	op   token.Token
}

func (err UnsupportedBinOP) Error() string {
	return fmt.Sprintf(
		"unsupported operation `%s` between %s and %s",
		err.op.String(), Type(err.a).String(), Type(err.b).String(),
	)
}

type Number struct {
	Num *big.Rat
	Fmt NumberFormat
}

func (a Number) String() string {
	return formatNum(a.Num, a.Fmt, "", 2)
}

func (a Number) BinOP(op token.Token, b Value) (Value, error) {
	// TODO: int type casts lose precision
	if op == token.TO {
		return nil, errors.New("unitless conversion")
	}
	switch b.(type) {
	case Number:
		b := b.(Number)
		var n *big.Rat
		switch op {
		case token.ADD:
			n = new(big.Rat).Add(a.Num, b.Num)
		case token.SUB:
			n = new(big.Rat).Sub(a.Num, b.Num)
		case token.MUL:
			n = new(big.Rat).Mul(a.Num, b.Num)
		case token.QUO:
			n = new(big.Rat).Quo(a.Num, b.Num)
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
			n = new(big.Rat).Set(a.Num)
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
		case token.ADD, token.SUB:
			if b.Units.Dimension().IsPure() {
				num := new(big.Rat).Set(a.Num)
				num.Mul(num, unit.Normalize(b.Num, b.Units))
				return a.BinOP(op, Number{Num: num})
			}
		}
	}
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (a Number) UnOP(op token.Token) (Value, error) {
	if op == token.SUB {
		return Number{Num: new(big.Rat).Neg(a.Num)}, nil
	}
	return nil, errors.New("unsupported unary operation: " + op.String())
}

func (a Number) In(fmt string) (Value, error) {
	f, ok := StringToFormat[fmt]
	if !ok {
		return nil, errors.New("unrecognized number format: " + fmt)
	}
	return Number{Num: a.Num, Fmt: f}, nil
}
