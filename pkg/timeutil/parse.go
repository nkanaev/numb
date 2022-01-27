package timeutil

import (
	"errors"
	"time"
)

var dateLayouts = []string{
	"2006-01-02",
	"2006/01/02",
	"2006.01.02",

	"02 Jan 2006",
	"_2 Jan 2006",
	"02 January 2006",
	"_2 January 2006",

	"01/02/2006",
	"01.02.2006",

	"2006-01-02T15:04",
	"2006/01/02 15:04",

	"02 Jan 2006 15:04",
	"_2 Jan 2006 15:04",
	"02 January 2006 15:04",
	"_2 January 2006 15:04",
}

var thisYearLayouts = []string{
	"2 Jan",
	"Jan 2",
}

var unknownFormat = errors.New("unknown time format")

func parseLocal(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

func Parse(value string) (time.Time, error) {
	for _, layout := range dateLayouts {
		if t, err := parseLocal(layout, value); err == nil {
			return t, nil
		}
	}

	for _, layout := range thisYearLayouts {
		if t, err := parseLocal(layout, value); err == nil {
			t = t.AddDate(time.Now().Year(), 0, 0)
			return t, nil
		}
	}

	return time.Time{}, unknownFormat
}
