package unit

import (
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/dimension"
)

type Unit struct {
	name   string
	value  *big.Rat
	offset *big.Rat
	dim    dimension.Dimension
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

func (bu unit) Expand() map[string]*Unit {
	names := splitlist(bu.names)

	result := make(map[string]*Unit)
	unit := &Unit{
		name:   names[0],
		value:  bu.value,
		offset: bu.offset,
		dim:    bu.u,
	}

	for _, alias := range names {
		result[alias] = unit
	}

	return result
}

var db = map[string]*Unit{}

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
		return u
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
			return u
		}
	}
	return nil
}

func init() {
	for _, bu := range units {
		for key, val := range bu.Expand() {
			db[key] = val
		}
	}
}

func Add(name string, num *big.Rat, unit UnitList) {
	db[name] = &Unit{
		name:  name,
		value: unit.normalize(num),
		dim:   unit.Dimension(),
	}
}
