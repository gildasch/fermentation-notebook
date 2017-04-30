package model

import (
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
