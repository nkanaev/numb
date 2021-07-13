package dimension

import "testing"

func TestMeasures(t *testing.T) {
	for m := Measure(0); m < end_measures; m += 1 {
		if m.String() == "" {
			t.Errorf("measure #%d is missing name", uint(m))
		}
		if m.Dim().Exp(2).Equals(m.Dim()) {
			t.Errorf("measure #%d is missing dimension", uint(m))
		}
	}
}
