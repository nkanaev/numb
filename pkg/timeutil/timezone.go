package timeutil

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"
)

func FindLocation(name string) (*time.Location, error) {
	if strings.EqualFold(name, "utc") {
		return time.UTC, nil
	}
	if strings.EqualFold(name, "local") {
		return time.Local, nil
	}

	var zoneDirs = []string{
		"/usr/share/zoneinfo",
		"/usr/share/lib/zoneinfo",
		"/usr/lib/locale/TZ",
	}
	for _, dir := range zoneDirs {
		if loc := findLocationFromDir(name, dir, ""); loc != nil {
			return loc, nil
		}
	}
	return nil, nil
}

func findLocationFromDir(name, rootdir, subdir string) *time.Location {
	files, _ := ioutil.ReadDir(rootdir + "/" + subdir)

	for _, f := range files {
		if !unicode.IsUpper(rune(f.Name()[0])) {
			continue
		}

		if f.IsDir() && len(subdir) == 0 {
			if loc := findLocationFromDir(name, rootdir, subdir+"/"+f.Name()); loc != nil {
				return loc
			}
		} else {
			tzname := strings.TrimLeft(subdir+"/"+f.Name(), "/")
			if tzNameMatches(name, tzname) {
				if tzdata, err := os.ReadFile(rootdir + "/" + tzname); err == nil {
					if tzinfo, err := time.LoadLocationFromTZData(tzname, tzdata); err == nil {
						return tzinfo
					}
				}
			}
		}
	}
	return nil
}

func normalizeTZName(name string) string {
	name = strings.ReplaceAll(name, "_", "")
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ToLower(name)
	return name
}

func tzNameMatches(name, tzname string) bool {
	name = normalizeTZName(name)
	tzname = normalizeTZName(tzname)

	if strings.ContainsRune(tzname, '/') && !strings.ContainsRune(name, '/') {
		chunks := strings.Split(tzname, "/")
		tzname := chunks[len(chunks)-1]
		if strings.EqualFold(name, tzname) {
			return true
		}
	}

	return strings.EqualFold(name, tzname)
}
