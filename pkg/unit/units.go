package unit

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/nkanaev/numb/pkg/dimension"
	r "github.com/nkanaev/numb/pkg/ratutils"
)

type namedUnit struct {
	name string
	unit Unit
	exp  int
}

type Units []namedUnit

func (u1 Units) Conforms(u2 Units) bool {
	return u1.Dimension().Equals(u2.Dimension())
}

func (u1 Units) IsSingle() bool {
	return len(u1) == 1 && u1[0].exp == 1
}

func (units Units) Simplify() Units {
	exps := make(map[dimension.Dimension]int)
	tmpunits := make(Units, 0)

	for _, nu := range units {
		key := nu.unit.dim
		if _, seen := exps[key]; seen {
			exps[key] += nu.exp
		} else {
			exps[key] += nu.exp
			tmpunits = append(tmpunits, nu)
		}
	}

	newunits := make(Units, 0)
	for _, nu := range tmpunits {
		key := nu.unit.dim
		if exp := exps[key]; exp != 0 {
			nu := namedUnit{name: nu.name, unit: nu.unit, exp: exp}
			newunits = append(newunits, nu)
		}
	}
	return newunits
}

func (u Units) Dimension() dimension.Dimension {
	var d dimension.Dimension
	for _, x := range u {
		d = d.Add(x.unit.dim.Exp(x.exp))
	}
	return d
}

func (units Units) normalize(n *big.Rat) *big.Rat {
	if len(units) == 0 {
		return n
	}
	num := new(big.Rat).Set(n)
	if units.IsSingle() {
		// (n + u.offset) * u.value
		u := units[0].unit
		if u.offset != nil {
			num.Add(num, u.offset)
		}
		num = num.Mul(num, u.value)
	} else {
		for _, u := range units {
			num.Mul(num, r.ExpInt(u.unit.value, u.exp))
		}
	}
	return num
}

func (units Units) denormalize(n *big.Rat) *big.Rat {
	if len(units) == 0 {
		return n
	}
	num := new(big.Rat).Set(n)
	if units.IsSingle() {
		// (n / u.value) - u.offset
		u := units[0].unit
		num.Quo(num, u.value)
		if u.offset != nil {
			num.Sub(num, u.offset)
		}
	} else {
		for _, u := range units {
			num.Quo(num, r.ExpInt(u.unit.value, u.exp))
		}
	}
	return num
}

func (u Units) String() string {
	if len(u) == 2 && u[0].exp == 1 && u[1].exp == -1 {
		return u[0].name + "/" + u[1].name
	}

	b := make([]string, 0, len(u))
	for _, entry := range u {
		if entry.exp == 0 {
			continue
		} else if entry.exp == 1 {
			b = append(b, entry.name)
		} else {
			b = append(b, fmt.Sprintf("%s^%d", entry.name, entry.exp))
		}
	}
	return strings.Join(b, " ")
}

func (u1 Units) Exp(x int) Units {
	u2 := Units{}
	for _, u := range u1 {
		u2 = append(u2, namedUnit{
			unit: u.unit,
			exp:  u.exp * x,
		})
	}
	return u2
}

func (this Units) Mul(other Units) Units {
	c := Units{}
	c = append(c, this...)
	c = append(c, other...)
	return c
}

func (this Units) Quo(other Units) Units {
	c := Units{}
	c = append(c, this...)
	for _, u := range other {
		c = append(c, namedUnit{
			unit: u.unit,
			exp:  -u.exp,
		})
	}
	return c
}
