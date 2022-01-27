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

var unknownFormat = errors.New("unknown time format")

func Parse(value string) (time.Time, error) {
	tz := time.Local
	for _, layout := range dateLayouts {
		if t, err := time.ParseInLocation(layout, value, tz); err == nil {
			return t, nil
		}
	}
	return time.Time{}, unknownFormat
}
