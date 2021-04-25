package unit

import (
	"reflect"
	"testing"
)

func TestUnitGet(t *testing.T) {
	for _, text := range []string{"m", "meter", "metre"} {
		have := Get(text)
		want := &Unit{
			name: "m",
			value: 1,
			dimension: LENGTH,
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("\ntext: %s\nwant: %s\nhave: %s", text, want, have)
		}
	}
}
