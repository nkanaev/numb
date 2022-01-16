package value

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"

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
	loc, err := getLocation(fmt)
	if loc != nil {
		return Time{ts: t.ts.In(loc), fmt: t.fmt}, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, errors.New("unknown timezone: " + fmt)
}

func (t Time) String() string {
	return t.ts.Format(t.fmt)
}

func getLocation(name string) (*time.Location, error) {
	if strings.EqualFold(name, "utc") {
		return time.UTC, nil
	}
	if strings.EqualFold(name, "local") {
		return time.Local, nil
	}

	var zoneDirs = []string{
		"/usr/share/zoneinfo",
		"/usr/share/lib/zoneinfo",
		"/usr/lib/locale/TZ",
	}
	for _, dir := range zoneDirs {
		if loc := findLocationFromDir(name, dir, ""); loc != nil {
			return loc, nil
		}
	}
	return nil, nil
}

func findLocationFromDir(name, rootdir, subdir string) *time.Location {
    files, _ := ioutil.ReadDir(rootdir + "/" + subdir)

    for _, f := range files {
        if !unicode.IsUpper(rune(f.Name()[0])) {
            continue
        }

        if f.IsDir() && len(subdir) == 0 {
			if loc := findLocationFromDir(name, rootdir, subdir + "/" + f.Name()); loc != nil {
				return loc
			}
        } else {
			tzname := strings.TrimLeft(subdir + "/" + f.Name(), "/")
			if tzNameMatches(name, tzname) {
				if tzdata, err := os.ReadFile(rootdir + "/" + tzname); err == nil {
					if tzinfo, err := time.LoadLocationFromTZData(tzname, tzdata); err == nil {
						return tzinfo
					}
				}
			}
        }
    }
	return nil
}

func normalizeTZName(name string) string {
	name = strings.ReplaceAll(name, "_", "")
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ToLower(name)
	return name
}

func tzNameMatches(name, tzname string) bool {
	name = normalizeTZName(name)
	tzname = normalizeTZName(tzname)

	if strings.ContainsRune(tzname, '/') && !strings.ContainsRune(name, '/') {
		chunks := strings.Split(tzname, "/")
		tzname := chunks[len(chunks)-1]
		if strings.EqualFold(name, tzname) {
			return true
		}
	}

	return strings.EqualFold(name, tzname)
}
