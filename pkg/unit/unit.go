package unit

import (
	"math/big"
	"strings"
)

type Dimension uint

const (
	LENGTH Dimension = 1 << iota
	TEMPERATURE
	AREA
	VOLUME
	MASS
	TIME
	ANGLE
	DIGITAL
	CURRENCY
)

type Unit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	dimension Dimension
}

type baseUnit struct {
	name         string
	value        *big.Rat
	offset       *big.Rat
	aliases      []string
	shortaliases []string
	dimension    Dimension
	prefixes     *[]prefix
	prefixpow    int
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
	for _, alias := range bu.shortaliases {
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
				name:      pr.abbr + bu.name,
				value:     prefixValue,
				offset:    bu.offset,
				dimension: bu.dimension,
			}

			result[pr.abbr+bu.name] = prefixUnit
			for _, alias := range bu.aliases {
				result[pr.name+alias] = prefixUnit
			}
			for _, alias := range bu.shortaliases {
				result[pr.abbr+alias] = prefixUnit
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
	if u, ok := db[x]; ok {
		return u
	}
	if u, ok := db[strings.ToLower(x)]; ok {
		return u
	}
	if u, ok := db[strings.ToUpper(x)]; ok {
		return u
	}
	return nil
}

func init() {
	unitLists := [][]baseUnit{
		lengthUnits,
		temperatureUnits,
		volumeUnits,
		areaUnits,
		timeUnits,
		digitalUnits,
		angleUnits,
		massUnits,
	}
	for _, unitList := range unitLists {
		for _, bu := range unitList {
			for key, val := range bu.Expand() {
				db[key] = val
			}
		}
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
			name: code,
			value: new(big.Rat).SetFloat64(1 / cur.Rate),
			dimension: CURRENCY,
		}
		db[code] = u
	}
}
