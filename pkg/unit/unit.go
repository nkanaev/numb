package unit

import (
	"fmt"
	"math/big"
	"strings"
)

type Dimension uint32

const (
	LENGTH Dimension = iota
	TEMPERATURE
	AREA
	VOLUME
	MASS
	TIME
	ANGLE
	DIGITAL
	CURRENCY
	FREQUENCY
	ELECTRIC_CURRENT
	LUMINOUS_INTENSITY
	AMOUNT_OF_SUBSTANCE
	POWER
	FORCE
	ENERGY
	ELECTRIC_CHARGE
	ELECTRIC_POTENTIAL
	ELECTRIC_CAPACITANCE
	ELECTRIC_CONDUCTANCE
	MAGNETIC_FLUX
	MAGNETIC_FLUX_DENSITY
	ELECTRIC_INDUCTANCE
	ELECTRIC_RESISTANCE
	PRESSURE
	RADIOACTIVITY
	SOLID_ANGLE
	IONIZING_RADIATION
	CATALYCTIC_ACTIVITY
	RADIATION_DOSE
	LUMINOUS_FLUX
	ILLUMINANCE
)

var dimensionNames = map[Dimension]string{
	LENGTH:                "length",
	TEMPERATURE:           "temperature",
	AREA:                  "area",
	VOLUME:                "volume",
	MASS:                  "mass",
	TIME:                  "time",
	ANGLE:                 "angle",
	DIGITAL:               "digital",
	CURRENCY:              "currency",
	FREQUENCY:             "frequency",
	ELECTRIC_CURRENT:      "electric current",
	LUMINOUS_INTENSITY:    "luminous intensity",
	AMOUNT_OF_SUBSTANCE:   "amount of substance",
	POWER:                 "power",
	FORCE:                 "force",
	ENERGY:                "energy",
	ELECTRIC_CHARGE:       "electric charge",
	ELECTRIC_POTENTIAL:    "electric potential",
	ELECTRIC_CAPACITANCE:  "electric capacitance",
	ELECTRIC_CONDUCTANCE:  "electric conductance",
	MAGNETIC_FLUX:         "magnetic flux",
	MAGNETIC_FLUX_DENSITY: "magnetic flux density",
	ELECTRIC_INDUCTANCE:   "electric inductance",
	ELECTRIC_RESISTANCE:   "electric resistance",
	PRESSURE:              "pressure",
	RADIOACTIVITY:         "radiactivity",
	SOLID_ANGLE:           "solid angle",
	IONIZING_RADIATION:    "ionizing radiation",
	CATALYCTIC_ACTIVITY:   "catalyctic activity",
	RADIATION_DOSE:        "radiation dose",
	LUMINOUS_FLUX:         "luminous flux",
	ILLUMINANCE:           "illuminance",
}

func (d Dimension) String() string {
	return dimensionNames[d]
}

type Unit struct {
	name      string
	value     *big.Rat
	offset    *big.Rat
	dimension Dimension
}

type baseUnit struct {
	d         Dimension
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

func (bu baseUnit) Expand() map[string]*Unit {
	shortforms := splitlist(bu.name)
	longforms := splitlist(bu.long)
	name := shortforms[0]

	result := make(map[string]*Unit, 0)
	unit := &Unit{
		name:      name,
		value:     bu.value,
		offset:    bu.offset,
		dimension: bu.d,
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
				dimension: bu.d,
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

func (u *Unit) String() string {
	return u.name
}

var db = map[string]*Unit{}

func Get(x string) *Unit {
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
	prevd := ILLUMINANCE
	for _, bu := range units {
		if bu.d != prevd {
			fmt.Println("")
			fmt.Println(bu.d)
			prevd = bu.d
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
			dimension: CURRENCY,
		}
		db[code] = u
	}
}
