package value

import "math/big"

type Value struct {
	Num *big.Rat
}

func NewInt(x int64) Value {
	return Value{Num: big.NewRat(x, 1)}
}

func (a Value) Mul(b Value) Value {
	return Value{Num: new(big.Rat).Mul(a.Num, b.Num)}
}

func (a Value) Add(b Value) Value {
	return Value{Num: new(big.Rat).Add(a.Num, b.Num)}
}

func (a Value) Sub(b Value) Value {
	return Value{Num: new(big.Rat).Sub(a.Num, b.Num)}
}

func (a Value) Quo(b Value) Value {
	return Value{Num: new(big.Rat).Quo(a.Num, b.Num)}
}

func (a Value) String() string {
	return a.Num.RatString()
}

func (a Value) Eval(map[string]Value) Value {
	return a
}
