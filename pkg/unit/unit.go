package unit

import (
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/unit/dimension"
)

type Unit struct {
	name    string
	value   *big.Rat
	offset  *big.Rat
	measure dimension.Measure
	// TODO: prefix value
}

func (u Unit) String() string {
	return u.name
}

type unitDef struct {
	u         dimension.Measure
	name      string
	long      string
	value     *big.Rat
	offset    *big.Rat
	prefixes  *[]prefix
	prefixpow int
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

func (bu unitDef) Expand() map[string]*Unit {
	shortforms := splitlist(bu.name)
	longforms := splitlist(bu.long)
	name := shortforms[0]

	result := make(map[string]*Unit)
	unit := &Unit{
		name:    name,
		value:   bu.value,
		offset:  bu.offset,
		measure: bu.u,
	}

	for _, alias := range shortforms {
		result[alias] = unit
	}
	for _, alias := range longforms {
		result[alias] = unit
	}

	if bu.prefixes != nil {
		for _, pr := range *bu.prefixes {
			prefixValue := big.NewRat(1, 1).Set(pr.value)
			if bu.prefixpow > 0 {
				x := new(big.Rat).Set(prefixValue)
				for i := 1; i < bu.prefixpow; i++ {
					prefixValue.Mul(prefixValue, x)
				}
			}
			prefixValue.Mul(prefixValue, bu.value)
			prefixUnit := &Unit{
				name:    pr.abbr + name,
				value:   prefixValue,
				offset:  bu.offset,
				measure: bu.u,
			}

			for _, alias := range longforms {
				result[pr.name+alias] = prefixUnit
			}
			for _, alias := range shortforms {
				result[pr.abbr+alias] = prefixUnit
			}
		}
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
		return nil, false
	}
	return UnitList{unitEntry{Unit: *u, Exp: 1}}, true
}

func getNamedUnit(x string) *Unit {
	if u, ok := db[x]; ok {
		return u
	}
	for _, suffix := range []string{"s", "es"} {
		if u, ok := db[strings.TrimSuffix(x, suffix)]; ok {
			return u
		}
	}
	if strings.HasSuffix(x, "ies") {
		if u, ok := db[strings.TrimSuffix(x, "ies") + "y"]; ok {
			return u
		}
	}
	for _, prefix := range metricPrefixes {
		if strings.HasPrefix(x, prefix.name) {
			if u, ok := db[strings.TrimPrefix(x, prefix.name)]; ok && u.offset == nil {
				return &Unit{
					name: x,
					value: new(big.Rat).Mul(u.value, prefix.value),
					measure: u.measure,
				}
			}
			break
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
	measure, found := unit.Dimension().Measure()
	if !found {
		panic("cannot create unit of unknown measure: " + unit.String())
	}
	db[name] = &Unit{
		name:    name,
		value:   unit.normalize(num),
		measure: measure,
	}
}
