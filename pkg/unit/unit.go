package unit

import (
	"math/big"
	"strings"

	d "github.com/nkanaev/numb/pkg/dimension"
	r "github.com/nkanaev/numb/pkg/ratutils"
)

type Unit struct {
	value  *big.Rat
	offset *big.Rat
	dim    d.Dimension
}

var db = map[string]Unit{}

func Must(x string) Units {
	u, ok := Get(x)
	if !ok {
		panic("invalid unit: " + x)
	}
	return u
}

func Get(x string) (Units, bool) {
	u := getNamedUnit(x)
	if u == nil {
		suffixes := map[string]string{
			"s":   "",
			"es":  "",
			"ies": "y",
		}
		for suffix, substitute := range suffixes {
			if u = getNamedUnit(strings.TrimSuffix(x, suffix) + substitute); u != nil {
				break
			}
		}
		if u == nil {
			return nil, false
		}
	}
	return Units{namedUnit{name: x, unit: *u, exp: 1}}, true
}

func getNamedUnit(x string) *Unit {
	if u, ok := db[x]; ok {
		return &u
	}
	for prefix, prefixVal := range prefixes {
		if u, ok := db[strings.TrimPrefix(x, prefix)]; ok && u.offset == nil {
			return &Unit{
				value: new(big.Rat).Mul(u.value, prefixVal),
				dim:   u.dim,
			}
		}
	}
	for name, u := range db {
		if strings.EqualFold(name, x) {
			return &u
		}
	}
	return nil
}

func init() {
	for key, val := range Defaults() {
		db[key] = val
	}
}

func Add(name string, num *big.Rat, unit Units) {
	db[name] = Unit{
		value: unit.normalize(num),
		dim:   unit.Dimension(),
	}
}

func Defaults() map[string]Unit {
	defaults := make(map[string]Unit)

	baseunits := map[string]d.Dimension{
		"LENGTH":              d.LENGTH,
		"MASS":                d.MASS,
		"TIME":                d.TIME,
		"AMOUNT_OF_SUBSTANCE": d.AMOUNT_OF_SUBSTANCE,
		"DIGITAL":             d.DIGITAL,
		"TEMPERATURE":         d.TEMPERATURE,
		"CURRENCY":            d.CURRENCY,
		"ELECTRIC_CURRENT":    d.ELECTRIC_CURRENT,
		"LUMINOUS_INTENSITY":  d.LUMINOUS_INTENSITY,
	}
	for name, dim := range baseunits {
		defaults[name] = Unit{value: r.ONE, dim: dim}
	}

	// non-linear units
	for _, name := range []string{"degC", "celsius"} {
		defaults[name] = Unit{value: r.ONE, offset: r.Must("273.15"), dim: d.TEMPERATURE}
	}
	for _, name := range []string{"degF", "fahrenheit"} {
		defaults[name] = Unit{value: big.NewRat(10, 18), offset: r.Must("459.67"), dim: d.TEMPERATURE}
	}

	return defaults
}

var prefixes = map[string]*big.Rat{
	// SI
	"d":     r.Exp(10, -1),
	"deci":  r.Exp(10, -1),
	"c":     r.Exp(10, -2),
	"centi": r.Exp(10, -2),
	"m":     r.Exp(10, -3),
	"milli": r.Exp(10, -3),
	"u":     r.Exp(10, -6),
	"micro": r.Exp(10, -6),
	"n":     r.Exp(10, -9),
	"nano":  r.Exp(10, -9),
	"p":     r.Exp(10, -12),
	"pico":  r.Exp(10, -12),
	"f":     r.Exp(10, -15),
	"femto": r.Exp(10, -15),
	"a":     r.Exp(10, -18),
	"atto":  r.Exp(10, -18),
	"z":     r.Exp(10, -21),
	"zepto": r.Exp(10, -21),
	"y":     r.Exp(10, -24),
	"yocto": r.Exp(10, -24),
	"da":     r.Exp(10, 1),
	"deca":  r.Exp(10, 1),
	"h":     r.Exp(10, 2),
	"hecto": r.Exp(10, 2),
	"k":     r.Exp(10, 3),
	"kilo":  r.Exp(10, 3),
	"M":     r.Exp(10, 6),
	"mega":  r.Exp(10, 6),
	"G":     r.Exp(10, 9),
	"giga":  r.Exp(10, 9),
	"T":     r.Exp(10, 12),
	"tera":  r.Exp(10, 12),
	"P":     r.Exp(10, 15),
	"peta":  r.Exp(10, 15),
	"E":     r.Exp(10, 18),
	"exa":   r.Exp(10, 18),
	"Z":     r.Exp(10, 21),
	"zetta": r.Exp(10, 21),
	"Y":     r.Exp(10, 24),
	"yotta": r.Exp(10, 24),
	// digital
	"Ki":    r.Exp(1024, 1),
	"kibi":  r.Exp(1024, 1),
	"Mi":    r.Exp(1024, 2),
	"mebi":  r.Exp(1024, 2),
	"Gi":    r.Exp(1024, 3),
	"gibi":  r.Exp(1024, 3),
	"Ti":    r.Exp(1024, 4),
	"tebi":  r.Exp(1024, 4),
	"Pi":    r.Exp(1024, 5),
	"pebi":  r.Exp(1024, 5),
	"Ei":    r.Exp(1024, 6),
	"exi":   r.Exp(1024, 6),
	"Zi":    r.Exp(1024, 7),
	"zebi":  r.Exp(1024, 7),
	"Yi":    r.Exp(1024, 8),
	"yobi":  r.Exp(1024, 8),
}
