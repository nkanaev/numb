package value

import (
	"errors"
	"math/big"

	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
)

type Percent struct {
	Num *big.Rat
}

func (a Percent) BinOP(op token.Token, b Value) (Value, error) {
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (p Percent) UnOP(op token.Token) (Value, error) {
	return nil, errors.New("unsupported unary operation: " + op.String())
}

func (p Percent) In(fmt string) (Value, error) {
	return nil, errors.New("percentages cannot be formatted")
}

func (p Percent) String() string {
	if p.Num.IsInt() {
		return p.Num.Num().String() + "%"
	}
	return p.Num.String() + "%"
}

func (p Percent) Apply(num *big.Rat) *big.Rat {
	num = new(big.Rat).Set(num)
	num.Mul(num, p.Num)
	num.Quo(num, ratutils.HUNDRED)
	return num
}
