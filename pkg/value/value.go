package value

import (
	"math/big"
	"time"

	"github.com/nkanaev/numb/pkg/token"
)

type ValueType int

const (
	TYPE_UNKNOWN ValueType = iota
	TYPE_NUMBER
	TYPE_UNIT
	TYPE_TIME
	TYPE_PERCENT
)

func (t ValueType) String() string {
	switch t {
	case TYPE_NUMBER:
		return "number"
	case TYPE_UNIT:
		return "unit"
	case TYPE_TIME:
		return "time"
	case TYPE_PERCENT:
		return "percent"
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
	case Time:
		return TYPE_TIME
	case Percent:
		return TYPE_PERCENT
	default:
		return TYPE_UNKNOWN
	}
}

type Value interface {
	BinOP(token.Token, Value) (Value, error)
	UnOP(token.Token) (Value, error)
	In(string) (Value, error)
	String() string
}

func Int64(x int64) Number {
	return Number{Num: new(big.Rat).SetInt64(x)}
}

func Float64(x float64) Number {
	return Number{Num: new(big.Rat).SetFloat64(x)}
}

func GetNamedTime(name string) *Time {
	switch name {
	case "now":
		return &Time{ts: time.Now(), fmt: time.ANSIC}
	case "time":
		return &Time{ts: time.Now(), fmt: "15:04"}
	}
	return nil
}
