package unit

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/unit/dimension"
)

type Unit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	// TODO: rename to measure
	dimension dimension.Measure
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
	info      string
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
		name:      name,
		value:     bu.value,
		offset:    bu.offset,
		dimension: bu.u,
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
			prefixUnit := &Unit{
				name:      pr.abbr + name,
				value:     prefixValue,
				offset:    bu.offset,
				dimension: bu.u,
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
	if u, ok := db[strings.ToUpper(x)]; ok {
		return u
	}
	if u, ok := db[strings.ToLower(x)]; ok {
		return u
	}
	if u, ok := db[strings.TrimSuffix(x, "s")]; ok {
		return u
	}
	if u, ok := db[strings.TrimSuffix(x, "es")]; ok {
		return u
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

func Help() {
	prevd := dimension.ILLUMINANCE
	for _, bu := range units {
		if bu.u != prevd {
			fmt.Println("")
			fmt.Println(bu.u)
			prevd = bu.u
		}
		names := splitlist(bu.name)

		longforms := splitlist(bu.long)
		if len(longforms) > 0 {
			// prevent long/short forms from appearing twice (ex.: bar, ohm)
			if longforms[0] != names[0] {
				names = append(names, longforms...)
			}
		}
		fmt.Printf("    %-20s\n", strings.Join(names, ", "))
	}
}

type Currency struct {
	Code string
	Rate float64
}

func AddExchangeRates(currencies []Currency) {
	for _, cur := range currencies {
		code := strings.ToUpper(cur.Code)
		u := &Unit{
			name:      code,
			value:     new(big.Rat).SetFloat64(1 / cur.Rate),
			dimension: dimension.CURRENCY,
		}
		db[code] = u
	}
}
