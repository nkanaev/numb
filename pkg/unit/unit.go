package unit

import (
	"math/big"
)

type Dimension uint

const (
	LENGTH Dimension = 1 << iota
	TEMPERATURE
	VOLUME
	MASS
	TIME
	ANGLE
	DIGITAL
)

type Unit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	dimension Dimension
}

type baseUnit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	aliases   []string
	dimension Dimension
	prefixes  *[]prefix
}

func (bu baseUnit) Expand() map[string]*Unit {
	result := make(map[string]*Unit, 0)
	unit := &Unit{
		name:      bu.name,
		value:     bu.value,
		offset:    bu.offset,
		dimension: bu.dimension,
	}

	result[bu.name] = unit
	for _, alias := range bu.aliases {
		result[alias] = unit
	}

	if bu.prefixes != nil {
		for _, pr := range *bu.prefixes {
			prefixValue := big.NewRat(1, 1).Set(pr.value)
			prefixValue.Mul(prefixValue, bu.value)
			prefixUnit := &Unit{
				name:      pr.abbr + bu.name,
				value:     prefixValue,
				offset:    bu.offset,
				dimension: bu.dimension,
			}

			result[pr.abbr+bu.name] = prefixUnit
			for _, alias := range bu.aliases {
				result[pr.name+alias] = prefixUnit
			}
		}
	}
	return result
}

func (u *Unit) String() string {
	return u.name
}

var db = map[string]*Unit{}

func Get(x string) *Unit {
	return db[x]
}

func init() {
	unitLists := [][]baseUnit{
		lengthUnits,
	}
	for _, unitList := range unitLists {
		for _, bu := range unitList {
			for key, val := range bu.Expand() {
				db[key] = val
			}
		}
	}
}
