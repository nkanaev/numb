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
	dimension dimension.NamedDimension
}

func (u Unit) String() string {
	return u.name
}

type unitEntry struct {
	Unit Unit
	Exp  int
}

type UnitList []unitEntry

func (u1 UnitList) Conforms(u2 UnitList) bool {
	d1 := u1.Dimension()
	d2 := u2.Dimension()
	return d1.Equals(d2)
}

func (u1 UnitList) Exp(x int) UnitList {
	u2 := UnitList{}
	for _, u := range u1 {
		u2 = append(u2, unitEntry{
			Unit: u.Unit,
			Exp: u.Exp * x,
		})
	}
	return u2
}

func (u UnitList) Dimension() dimension.Dimensions {
	var d dimension.Dimensions
	for _, x := range u {
		d = d.Add(x.Unit.dimension.Dims.Exp(x.Exp))
	}
	return d
}

func (u UnitList) String() string {
	if len(u) == 2 && u[0].Exp == 1 && u[1].Exp == -1 {
		return u[0].Unit.String() + "/" + u[1].Unit.String()
	}

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

func (this UnitList) Mul(other UnitList) UnitList {
	c := UnitList{}
	for _, u := range this {
		c = append(c, u)
	}
	for _, u := range other {
		c = append(c, u)
	}
	return c
}

func (this UnitList) Quo(other UnitList) UnitList {
	c := UnitList{}
	for _, u := range this {
		c = append(c, u)
	}
	for _, u := range other {
		c = append(c, unitEntry{
			Unit: u.Unit,
			Exp:  -u.Exp,
		})
	}
	return c
}

type unitDef struct {
	u         dimension.NamedDimension
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
			prefixValue.Mul(prefixValue, bu.value)
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
