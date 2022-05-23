package timeutil

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var dateLayouts = []string{
	"2006-01-02",
	"2006-1-2",
	"02-Jan-2006",
	"_2-Jan-2006",
	"02-January-2006",
	"_2-January-2006",
    "2006-Jan-02",
    "2006-Jan-_2",
    "2006-January-02",
    "2006-January-_2",
	"02-01-2006",
	"2-1-2006",

	"2006-01-02T15:04",
	"02-Jan-2006T15:04",
	"_2-Jan-2006T15:04",
	"02-January-2006T15:04",
	"_2-January-2006T15:04",
}

var (
    dateSeparators = []string{"-", "/", ".", " "}
    timeSeparators = []string{"T", " "}
    defaultDateSep = "-"
    defaultTimeSep = "T"
)

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
        for _, dateSep := range dateSeparators {
            for _, timeSep := range timeSeparators {
                layout := strings.ReplaceAll(layout, defaultDateSep, dateSep)
                layout = strings.ReplaceAll(layout, defaultTimeSep, timeSep)
                if t, err := parseLocal(layout, value); err == nil {
                    return t, nil
                }
            }
        }
	}

	for _, layout := range thisYearLayouts {
		if t, err := parseLocal(layout, value); err == nil {
			t = t.AddDate(time.Now().Year(), 0, 0)
			return t, nil
		}
	}

	// If value consists only of digits, assume it's timestamp
	if strings.Trim(value, "0123456789") == "" {
		switch {
		case len(value) <= 11:
			// min: 1970-01-01T01:00:00 (0 s)
			// max: 5138-11-16T09:46:39 (99999999999 s)
			num, _ := strconv.ParseInt(value, 10, 64)
			return time.Unix(num, 0), nil
		case len(value) <= 14:
			// min: 1973-03-03T09:46:40 (100000000000 ms -> 100000000 s)
			// max: 5138-11-16T09:46:39 (99999999999999 ms -> 99999999999 s)
			num, _ := strconv.ParseInt(value, 10, 64)
			return time.UnixMilli(num), nil
		case len(value) <= 17:
			// min: 1973-03-03T09:46:40 (100000000000000 us -> 100000000 s)
			// max: 5138-11-16T09:46:39 (100000000000000 us -> 100000000 s)
			num, _ := strconv.ParseInt(value, 10, 64)
			return time.UnixMicro(num), nil
		default:
			return time.Time{}, errors.New("timestamp value out of range")
		}
	}

	return time.Time{}, unknownFormat
}
