package unit

import (
	"math/big"
	"strings"

	d "github.com/nkanaev/numb/pkg/dimension"
	r "github.com/nkanaev/numb/pkg/ratutils"
)

type unitDef struct {
	value  *big.Rat
	offset *big.Rat
	dim    d.Dimension
}

var db = map[string]unitDef{}

func Must(x string) Units {
	u, ok := Get(x)
	if !ok {
		panic("invalid unit: " + x)
	}
	return u
}

func (u unitDef) ToUnits(name string) Units {
	return Units{unitEntry{name: name, unit: u, exp: 1}}
}

func Get(name string) (Units, bool) {
	return resolveUnit(name, true)
}

func resolveUnit(name string, subsuffix bool) (Units, bool) {
	// 1. find the exact definition
	if def, ok := db[name]; ok {
		return def.ToUnits(name), true
	}

	// 2. substitute suffixes
	if subsuffix {
		suffixes := map[string]string{
			"s":   "",
			"es":  "",
			"ies": "y",
		}
		for suffix, substitute := range suffixes {
			newname := strings.TrimSuffix(name, suffix) + substitute
			if u, ok := resolveUnit(newname, false); ok {
				return u, ok
			}
		}
	}

	// 3. search with prefix
	for prefix, prefixVal := range prefixes {
		if strings.HasPrefix(name, prefix) {
			newname := strings.TrimPrefix(name, prefix)
			if def, ok := db[name]; ok && def.offset == nil {
				newval := new(big.Rat).Mul(def.value, prefixVal)
				newdef := unitDef{value: newval, dim: def.dim}
				return newdef.ToUnits(newname), true
			}
		}
	}

	// 4. if all else fails, try case-insensitive match
	for unitname, unitdef := range db {
		if strings.EqualFold(name, unitname) {
			return unitdef.ToUnits(unitname), true
		}
	}
	for prefix, prefixVal := range prefixes {
		for unitname, unitdef := range db {
			// no fractional digital units
			if unitdef.dim.Equals(d.DIGITAL) && prefixVal.Cmp(r.ONE) < 0 {
				continue
			}

			newname := prefix + unitname
			if strings.EqualFold(name, newname) {
				newval := new(big.Rat).Mul(unitdef.value, prefixVal)
				newdef := unitDef{value: newval, dim: unitdef.dim}
				return newdef.ToUnits(newname), true
			}
		}
	}

	return nil, false
}

func init() {
	for key, val := range Defaults() {
		db[key] = val
	}
}

func Add(name string, num *big.Rat, unit Units) {
	db[name] = unitDef{
		value: unit.normalize(num),
		dim:   unit.Dimension(),
	}
}

func Defaults() map[string]unitDef {
	defaults := make(map[string]unitDef)

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
		defaults[name] = unitDef{value: r.ONE, dim: dim}
	}

	// non-linear units
	for _, name := range []string{"degC", "celsius"} {
		defaults[name] = unitDef{value: r.ONE, offset: r.Must("273.15"), dim: d.TEMPERATURE}
	}
	for _, name := range []string{"degF", "fahrenheit"} {
		defaults[name] = unitDef{value: big.NewRat(10, 18), offset: r.Must("459.67"), dim: d.TEMPERATURE}
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
	"da":    r.Exp(10, 1),
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
	"Ki":   r.Exp(1024, 1),
	"kibi": r.Exp(1024, 1),
	"Mi":   r.Exp(1024, 2),
	"mebi": r.Exp(1024, 2),
	"Gi":   r.Exp(1024, 3),
	"gibi": r.Exp(1024, 3),
	"Ti":   r.Exp(1024, 4),
	"tebi": r.Exp(1024, 4),
	"Pi":   r.Exp(1024, 5),
	"pebi": r.Exp(1024, 5),
	"Ei":   r.Exp(1024, 6),
	"exi":  r.Exp(1024, 6),
	"Zi":   r.Exp(1024, 7),
	"zebi": r.Exp(1024, 7),
	"Yi":   r.Exp(1024, 8),
	"yobi": r.Exp(1024, 8),
}
