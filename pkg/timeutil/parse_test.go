package timeutil

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	year := time.Now().Year()

	tz := time.Local
	tcases := []struct{
		Date string
		Want time.Time
	}{
		// YYYY MM DD
		{"2019-05-06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"2019/05/06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"2019.05.06", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},

		// DD MM YYYY
		{"05/06/2019", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"05.06.2019", time.Date(2019, time.May, 6, 0, 0, 0, 0, tz)},
		{"06 Jun 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},
		{"6 Jun 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},
		{"06 June 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},
		{"6 June 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},
		{"06 june 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},
		{"6 june 2019", time.Date(2019, time.June, 6, 0, 0, 0, 0, tz)},

		// YYYY MM DD HH MM
		{"2019-05-06T13:30", time.Date(2019, time.May, 6, 13, 30, 0, 0, tz)},
		{"2019/05/06 13:30", time.Date(2019, time.May, 6, 13, 30, 0, 0, tz)},
		{"06 Jun 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},
		{"6 Jun 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},
		{"06 June 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},
		{"6 June 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},
		{"06 june 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},
		{"6 june 2019 13:30", time.Date(2019, time.June, 6, 13, 30, 0, 0, tz)},

		// MM DD
		{"may 6", time.Date(year, time.May, 6, 0, 0, 0, 0, tz)},
		{"6 may", time.Date(year, time.May, 6, 0, 0, 0, 0, tz)},
	}

	for _, tcase := range tcases {
		date := tcase.Date
		want := tcase.Want
		have, err := Parse(date)
		if err != nil {
			t.Errorf("failed for `%s` (%s)", date, err)
		}
		if !want.Equal(have) {
			t.Errorf(
				"invalid date\ndate: %s\nwant: %q\nhave: %q",
				date, want, have,
			)
		}
	}
}
