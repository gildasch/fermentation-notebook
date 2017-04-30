package model

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseDuration adds the support of Days for time.ParseDuration
func ParseDuration(s string) (time.Duration, error) {
	if !strings.Contains(s, "d") {
		return time.ParseDuration(s)
	}

	var ret time.Duration
	dayAndRest := strings.Split(s, "d")
	nDay, err := strconv.ParseFloat(dayAndRest[0], 64)
	if err != nil {
		return ret, err
	}

	if dayAndRest[1] != "" {
		ret, err = time.ParseDuration(dayAndRest[1])
		if err != nil {
			return ret, err
		}
	}

	return ret + time.Duration(24*nDay)*time.Hour, nil
}

// DurationToString adds the support for Days for time.Duration.String()
func DurationToString(d time.Duration) string {
	// Handle sign
	sign := ""
	if d < 0 {
		d = -d
		sign = "-"
	}

	// Split days and rest
	days := int64(d) / (24 * int64(time.Hour))
	daysStr := strconv.FormatInt(days, 10)

	rest := d % (24 * time.Hour)
	// Crop after second
	rest = rest - (rest % time.Second)

	ret := sign + daysStr + "d" + rest.String()

	// Remove zeros
	re := regexp.MustCompile("([^0-9])0[dhms]")
	ret = re.ReplaceAllString(ret, "$1")

	return ret
}
