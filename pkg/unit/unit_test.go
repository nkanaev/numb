package unit

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/unit/dimension"
)

func TestUnitGet(t *testing.T) {
	for _, text := range []string{"m", "meter", "metre"} {
		have := Must(text)
		want := UnitList{unitEntry{
			Unit: Unit{
				name:    "m",
				value:   big.NewRat(1, 1),
				measure: dimension.LENGTH,
			},
			Exp: 1,
		}}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %s\nhave: %s", text, want, have)
		}
	}
}

func TestUnitGetPrefixed(t *testing.T) {
	for _, text := range []string{"km", "kilometer", "kilometre"} {
		have := Must(text)
		want := UnitList{unitEntry{
			Unit: Unit{
				name:    "km",
				value:   big.NewRat(1000, 1),
				measure: dimension.LENGTH,
			},
			Exp: 1,
		}}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %#v\nhave: %#v", text, want, have)
		}
	}

	for _, text := range []string{"cm", "centimeter", "centimetre"} {
		have := Must(text)
		want := UnitList{unitEntry{
			Unit: Unit{
				name:    "cm",
				value:   big.NewRat(1, 100),
				measure: dimension.LENGTH,
			},
			Exp: 1,
		}}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %#v (%s)\nhave: %#v (%s)", text, want, want, have, have)
		}
	}
}
