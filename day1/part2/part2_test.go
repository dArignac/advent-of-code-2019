package part2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuelWithRecursion(t *testing.T) {
	assert.Equal(t, 2, CalculateRequiredFuel(12))
	assert.Equal(t, 966, CalculateRequiredFuel(1969))
	assert.Equal(t, 50346, CalculateRequiredFuel(100756))
}

func TestCalculateRequiredFuelForMultipleMasses(t *testing.T) {
	assert.Equal(t, 2+966+50346, CalculateRequiredFuelForMultipleMasses([]int{12, 1969, 100756}))
}
