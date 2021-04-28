package unit

import (
	"math/big"
	"reflect"
	"testing"
)

func TestUnitGet(t *testing.T) {
	for _, text := range []string{"m", "meter", "metre"} {
		have := Get(text)
		want := &Unit{
			name:      "m",
			value:     big.NewRat(1, 1),
			dimension: LENGTH,
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %s\nhave: %s", text, want, have)
		}
	}
}

func TestUnitGetPrefixed(t *testing.T) {
	for _, text := range []string{"km", "kilometer", "kilometre"} {
		have := Get(text)
		want := &Unit{
			name:      "km",
			value:     big.NewRat(1000, 1),
			dimension: LENGTH,
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %#v\nhave: %#v", text, want, have)
		}
	}

	for _, text := range []string{"cm", "centimeter", "centimetre"} {
		have := Get(text)
		want := &Unit{
			name:      "cm",
			value:     big.NewRat(1, 100),
			dimension: LENGTH,
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %#v (%s)\nhave: %#v (%s)", text, want, want.value, have, have.value)
		}
	}
}
