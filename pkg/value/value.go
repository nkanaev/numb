package value

import "math/big"

type Value struct {
	Num *big.Rat
}

func NewInt(x int64) Value {
	return Value{Num: big.NewRat(x, 1)}
}

func Parse(x string) Value {
	num, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse number: " + x)
	}
	return Value{Num: num}
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

func (a Value) Lsh(b Value) Value {
	ia, ib := a.Num.Num(), uint(b.Num.Num().Uint64())
	num := big.NewRat(1, 1)
	num.Num().Lsh(ia, ib)
	return Value{Num: num}
}

func (a Value) Rsh(b Value) Value {
	ia, ib := a.Num.Num(), uint(b.Num.Num().Uint64())
	num := big.NewRat(1, 1)
	num.Num().Rsh(ia, ib)
	return Value{Num: num}
}

func (a Value) And(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().And(a.Num.Num(), b.Num.Num())
	return Value{Num: num}
}

func (a Value) Or(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Or(a.Num.Num(), b.Num.Num())
	return Value{Num: num}
}

func (a Value) Xor(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Xor(a.Num.Num(), b.Num.Num())
	return Value{Num: num}
}

func (a Value) Rem(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Rem(a.Num.Num(), b.Num.Num())
	return Value{Num: num}
}

func (a Value) Exp(b Value) Value {
	num := big.NewRat(1, 1)
	num.Num().Exp(a.Num.Num(), b.Num.Num(), nil)
	return Value{Num: num}
}

func (a Value) String() string {
	return a.Num.RatString()
}

func (a Value) Eval(map[string]Value) Value {
	return a
}
