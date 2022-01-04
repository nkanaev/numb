package value

import (
	"errors"
	"time"

	"github.com/nkanaev/numb/pkg/token"
)

type Time struct {
	ts time.Time
	fmt string
}

func (a Time) BinOP(op token.Token, b Value) (Value, error) {
	return nil, UnsupportedBinOP{a: a, b: b, op: op}
}

func (t Time) UnOP(op token.Token) (Value, error) {
	return nil, errors.New("unsupported unary operation: %s" + op.String())
}

func (t Time) In(fmt string) (Value, error) {
	// TODO: case-insensitive tz match
	if loc, _ := time.LoadLocation(fmt); loc != nil {
		return Time{ts: t.ts.In(loc), fmt: t.fmt}, nil
	}
	return nil, errors.New("unrecognized time format: " + fmt)
}

func (t Time) String() string {
	return t.ts.Format(t.fmt)
}
