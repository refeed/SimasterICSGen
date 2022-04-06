package simastertime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	expectedTime := time.Date(2022, 4, 8, 8, 15, 0, 0, time.FixedZone("UTC+7", 7*60*60))
	actualDate := Parse("8 April 2022", "08:15")

	assert.Equal(t, expectedTime, actualDate)
}

func TestParsePM(t *testing.T) {
	expectedTime := time.Date(2022, 4, 8, 19, 15, 0, 0, time.FixedZone("UTC+7", 7*60*60))
	actualDate := Parse("8 April 2022", "19:15")

	assert.Equal(t, expectedTime, actualDate)
}
