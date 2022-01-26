package timeutil

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tz := time.Local
	tcases := []struct{
		Line string
		Want time.Time
	}{
		// YYYY MM DD
		{"2019-05-06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"2019/05/06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"2019.05.06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},

		// DD MM YYYY
		{"05/06/2019", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"05.06.2019", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},

		// YYYY MM DD HH MM
		{"2019-05-06T13:30", time.Date(2019, time.May, 6, 13, 30, 0, 0, tz)},
		{"2019/05/06 13:30", time.Date(2019, time.May, 6, 13, 30, 0, 0, tz)},
	}

	for _, tcase := range tcases {
		line := tcase.Line
		want := tcase.Want
		have, err := Parse(line)
		if err != nil {
			t.Errorf("failed for `%s` (%s)", line, err)
		}
		if !want.Equal(have) {
			t.Errorf(
				"invalid date\nline: %s\nwant: %q\nhave: %q",
				line, want, have,
			)
		}
	}
}
