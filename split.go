package moment

import (
	"errors"
	"strconv"
	"strings"
)

func splitterFull(d string) map[string]int64 {
	fm := make(map[string]int64)

	hr := strings.Split(d, "h")[0]
	min := strings.Split(strings.Split(d, "m")[0], "h")[1]
	sec := strings.Split(strings.Split(d, "m")[1], "s")[0]

	h, err := strconv.ParseInt(hr, 0, 0)
	if err != nil {
		errors.New("Hours conversion error")
	}
	fm["hours"] = h

	m, err := strconv.ParseInt(min, 0, 0)
	if err != nil {
		errors.New("Minutes conversion erro")
	}
	fm["minutes"] = m

	s, err := strconv.ParseFloat(sec, 64)
	if err != nil {
		errors.New("Second conversion error")
	}
	fm["seconds"] = int64(s)

	return fm

}
func splitterWithoutHour(d string) map[string]int64 {
	fm := make(map[string]int64)
	min := strings.Split(d, "m")[0]
	sec := strings.Split(strings.Split(d, "m")[1], "s")[0]

	m, err := strconv.ParseInt(min, 0, 0)
	if err != nil {
		errors.New("Minutes conversion erro")
	}
	fm["minutes"] = m

	s, err := strconv.ParseFloat(sec, 64)
	if err != nil {
		errors.New("Second conversion error")
	}
	fm["seconds"] = int64(s)
	return fm
}
func splitterWithoutMinutes(sec string) map[string]int64 {
	fm := make(map[string]int64)
	s, err := strconv.ParseFloat(sec, 64)
	if err != nil {
		errors.New("Second conversion error")
	}
	fm["seconds"] = int64(s)
	return fm
}
