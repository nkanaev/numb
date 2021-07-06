package unit

import (
	"fmt"
	"math/big"
	"strings"
	"github.com/nkanaev/numb/pkg/unit/dimension"
)

// TODO: rename to Unit
type NamedUnit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	dimension dimension.Dimension
}

func (u NamedUnit) String() string {
	return u.name
}

type unitEntry struct {
	Unit NamedUnit
	Exp  int
}

// TODO: rename to UnitList
type Unit []unitEntry

func (u Unit) String() string {
	b := make([]string, 0, len(u))
	for _, entry := range u {
		if entry.Exp == 0 {
			continue
		} else if entry.Exp == 1 {
			b = append(b, entry.Unit.String())
		} else {
			b = append(b, fmt.Sprintf("%s^%d", entry.Unit.String(), entry.Exp))
		}
	}
	return strings.Join(b, " ")
}

func (this *Unit) Mul(other *Unit) *Unit {
	c := Unit{}
	for _, u := range *this {
		c = append(c, u)
	}
	for _, u := range *other {
		c = append(c, u)
	}
	return &c
}

func (this *Unit) Quo(other *Unit) *Unit {
	c := Unit{}
	for _, u := range *this {
		c = append(c, u)
	}
	for _, u := range *other {
		c = append(c, unitEntry{
			Unit: u.Unit,
			Exp:  -u.Exp,
		})
	}
	return &c
}

type unitDef struct {
	u         dimension.Dimension
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

func (bu unitDef) Expand() map[string]*NamedUnit {
	shortforms := splitlist(bu.name)
	longforms := splitlist(bu.long)
	name := shortforms[0]

	result := make(map[string]*NamedUnit)
	unit := &NamedUnit{
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
			prefixValue.Mul(prefixValue, bu.value)
			if bu.prefixpow > 0 {
				x := new(big.Rat).Set(prefixValue)
				for i := 1; i < bu.prefixpow; i++ {
					prefixValue.Mul(prefixValue, x)
				}
			}
			prefixUnit := &NamedUnit{
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

var db = map[string]*NamedUnit{}

func Get(x string) *Unit {
	u := getNamedUnit(x)
	if u == nil {
		return nil
	}
	return &Unit{unitEntry{Unit: *u, Exp: 1}}
}

func getNamedUnit(x string) *NamedUnit {
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
		u := &NamedUnit{
			name:      code,
			value:     new(big.Rat).SetFloat64(1 / cur.Rate),
			dimension: dimension.CURRENCY,
		}
		db[code] = u
	}
}
