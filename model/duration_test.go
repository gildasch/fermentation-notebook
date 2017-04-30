package model

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const NumberOfRands = 10

func randomDurations() (ret []time.Duration) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < NumberOfRands; i++ {
		ret = append(ret, time.Duration(r.Int63()))
	}
	return
}

func TestParseRegularDuration(t *testing.T) {
	durations := randomDurations()

	for _, d := range durations {
		actual, err := ParseDuration(d.String())

		assert.Nil(t, err)
		assert.Equal(t, d, actual)
	}
}

func TestParseDurationWithDays(t *testing.T) {
	durations := randomDurations()

	for _, d := range durations {
		str := d.String()
		expected := time.Duration(36*time.Hour) + d
		actual, err := ParseDuration("1.5d" + str)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestDurationToString(t *testing.T) {
	expected := "2d12h5m"
	d, err := ParseDuration(expected)
	assert.Nil(t, err)
	actual := DurationToString(d)

	assert.Equal(t, expected, actual)
}
