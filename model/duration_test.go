package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDuration(t *testing.T) {
	expected, _ := time.ParseDuration("36h")
	actual, err := ParseDuration("1.5d")

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
