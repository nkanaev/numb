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
	name         string
	value        *big.Rat
	offset       *big.Rat
	aliases      []string
	shortaliases []string
	dimension    Dimension
	prefixes     *[]prefix
	prefixpow    int
	description  string
	// TODO: kohm/kiloohm, kbar/kilobar, kilohm (vower omitted) edge cases
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

func getUnitList() [][]baseUnit {
	return [][]baseUnit{
		lengthUnits,
		temperatureUnits,
		volumeUnits,
		areaUnits,
		timeUnits,
		digitalUnits,
		angleUnits,
		massUnits,
		frequencyUnits,
		electricCurrentUnits,
		luminousIntensityUnits,
		amountOfSubstanceUnits,
		powerUnits,
		forceUnits,
		energyUnits,
		electricChargeUnits,
		electricPotentialUnits,
		electricCapaticanceUnits,
		electricConductanceUnits,
		magneticFluxUnits,
		magneticFluxDensityUnits,
		electricInductanceUnits,
		electricResistanceUnits,
		pressureUnits,
		radioactivityUnits,
		solidAngleUnits,
		ionizingRadiationUnits,
		catalycticActivityUnits,
		radiationDoseUnits,
	}
}

func init() {
	for _, unitList := range getUnitList() {
		for _, bu := range unitList {
			for key, val := range bu.Expand() {
				db[key] = val
			}
		}
	}
}

func Help() {
	for i, unitList := range getUnitList() {
		if i != 0 {
			fmt.Println("")
		}
		fmt.Println("#", unitList[0].dimension)
		for _, bu := range unitList {
			names := make([]string, 0, 1+len(bu.aliases)+len(bu.shortaliases))
			names = append(names, bu.name)
			for _, alias := range bu.shortaliases {
				names = append(names, alias)
			}
			for _, alias := range bu.aliases {
				names = append(names, alias)
			}

			var description string
			if bu.description != "" {
				description = " # " + bu.description
			}
			fmt.Printf("    %-16s%s\n", strings.Join(names, ", "), description)
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
			name:      code,
			value:     new(big.Rat).SetFloat64(1 / cur.Rate),
			dimension: CURRENCY,
		}
		db[code] = u
	}
}
