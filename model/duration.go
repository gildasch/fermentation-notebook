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
	nDay, err := strconv.Atoi(dayAndRest[0])
	if err != nil {
		return ret, err
	}
	d, err := time.ParseDuration(dayAndRest[1])
	if err != nil {
		return ret, err
	}

	return d + time.Duration(24*nDay)*time.Hour, nil
}
