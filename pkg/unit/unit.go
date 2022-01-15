package unit

import (
	"math/big"
	"strings"

	d "github.com/nkanaev/numb/pkg/dimension"
	r "github.com/nkanaev/numb/pkg/ratutils"
)

type Unit struct {
	name   string
	value  *big.Rat
	offset *big.Rat
	dim    d.Dimension
}

func (u Unit) String() string {
	return u.name
}

func splitlist(x string) []string {
	if len(x) == 0 {
		return nil
	}
	list := make([]string, 0)
	for _, chunk := range strings.Split(x, ",") {
		list = append(list, strings.TrimSpace(chunk))
	}
	return list
}

var db = map[string]Unit{}

func Must(x string) UnitList {
	u, ok := Get(x)
	if !ok {
		panic("invalid unit: " + x)
	}
	return u
}

func Get(x string) (UnitList, bool) {
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
	return UnitList{unitEntry{Unit: *u, Exp: 1}}, true
}

func getNamedUnit(x string) *Unit {
	if u, ok := db[x]; ok {
		return &u
	}
	for _, prefix := range prefixes {
		for _, name := range splitlist(prefix.names) {
			if strings.HasPrefix(x, name) {
				if u, ok := db[strings.TrimPrefix(x, name)]; ok && u.offset == nil {
					return &Unit{
						name:  x,
						value: new(big.Rat).Mul(u.value, prefix.value),
						dim:   u.dim,
					}
				}
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

func Add(name string, num *big.Rat, unit UnitList) {
	db[name] = Unit{
		name:  name,
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
		defaults[name] = Unit{name: name, value: r.ONE, dim: dim}
	}

	// non-linear units
	for _, name := range []string{"degC", "celsius"} {
		defaults[name] = Unit{
			name: name, value: r.ONE, offset: r.Must("273.15"),
			dim: d.TEMPERATURE,
		}
	}
	for _, name := range []string{"degF", "fahrenheit"} {
		defaults[name] = Unit{
			name: name, value: big.NewRat(10, 18), offset: r.Must("459.67"),
			dim: d.TEMPERATURE,
		}
	}

	return defaults
}
