package value

import (
	"errors"
	"fmt"
	"time"

	"github.com/nkanaev/numb/pkg/dimension"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type Time struct {
	ts  time.Time
	fmt string
}

func (a Time) BinOP(op token.Token, b Value) (Value, error) {
	if Type(b) == TYPE_UNIT {
		b := b.(Unit)
		if b.Units.Dimension() != dimension.TIME {
			return nil, fmt.Errorf("%s is not a measure of %s", b, dimension.TIME)
		}

		diff_nsec, err := unit.Convert(b.Num, b.Units, unit.Must("nanosecond"))
		if err != nil {
			return nil, err
		}
		if !diff_nsec.IsInt() {
			return nil, fmt.Errorf("cannot deal with fractional nanosecs")
		}
		diff_int64 := time.Duration(diff_nsec.Num().Int64())

		if op == token.ADD {
			return Time{ts: a.ts.Add(diff_int64), fmt: a.fmt}, nil
		} else if op == token.SUB {
			return Time{ts: a.ts.Add(-diff_int64), fmt: a.fmt}, nil
		}
	}
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
