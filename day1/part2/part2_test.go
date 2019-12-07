package part2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuelWithRecursion(t *testing.T) {
	assert.Equal(t, 2, CalculateRequiredFuelWithRecursion(12))
	assert.Equal(t, 966, CalculateRequiredFuelWithRecursion(1969))
	assert.Equal(t, 50346, CalculateRequiredFuelWithRecursion(100756))
}
