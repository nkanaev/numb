package value

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/nkanaev/numb/pkg/dimension"
	"github.com/nkanaev/numb/pkg/timeutil"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
)

type Time struct {
	ts  time.Time
	fmt string
}

func (a Time) BinOP(op token.Token, b Value) (Value, error) {
	switch b.(type) {
	case Time:
		b := b.(Time)
		switch op {
			case token.SUB:
				maxYears := 200
				ayear := a.ts.Year()
				byear := b.ts.Year()
				if math.Abs(float64(ayear-byear)) > float64(maxYears) {
					return nil, fmt.Errorf("time difference exceedes limitation (%d years)", maxYears)
				}
				ns_i64 := a.ts.Sub(b.ts).Nanoseconds()
				ns_rat := new(big.Rat).SetInt64(ns_i64)
				return Unit{Num: ns_rat, Units: unit.Must("nanosecond")}, nil
		}
	case Unit:
		b := b.(Unit)
		if b.Units.Dimension().Equals(dimension.TIME) {
			return nil, fmt.Errorf("%s is not a measure of time", b)
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
	loc, err := timeutil.FindLocation(fmt)
	if loc != nil {
		return Time{ts: t.ts.In(loc), fmt: t.fmt}, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unknown timezone: " + fmt)
}

func NewTime(t time.Time) Time {
	return Time{ts: t}
}

func (t Time) String() string {
	if t.fmt == "" {
		return t.ts.Format(time.RFC1123)
	}
	return t.ts.Format(t.fmt)
}

func GetNamedTime(name string) *Time {
	now := time.Now()

	datefmt := "02 Jan 2006"

	switch name {
	case "now":
		return &Time{ts: now, fmt: time.ANSIC}
	case "time":
		return &Time{ts: now, fmt: "15:04"}
	case "date", "today":
		return &Time{ts: now, fmt: datefmt}
	case "tomorrow":
		return &Time{ts: now.Add(time.Hour*24), fmt: datefmt}
	case "yesterday":
		return &Time{ts: now.Add(-time.Hour*24), fmt: datefmt}
	}

	return nil
}
