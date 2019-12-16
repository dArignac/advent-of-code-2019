package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuel(t *testing.T) {
	assert.Equal(t, 0, CalculateRequiredFuel(0))
	assert.Equal(t, 0, CalculateRequiredFuel(2))
	assert.Equal(t, 2, CalculateRequiredFuel(12))
	assert.Equal(t, 2, CalculateRequiredFuel(14))
	assert.Equal(t, 654, CalculateRequiredFuel(1969))
	assert.Equal(t, 33583, CalculateRequiredFuel(100756))
}

func TestCalculateRequiredFuelWithRecursion(t *testing.T) {
	assert.Equal(t, 2, CalculateRequiredFuelWithRecursion(12))
	assert.Equal(t, 966, CalculateRequiredFuelWithRecursion(1969))
	assert.Equal(t, 50346, CalculateRequiredFuelWithRecursion(100756))
}

func TestRunCalculationFunctionForMasses(t *testing.T) {
	sqrt := func(x int) int {
		return x * x
	}
	assert.Equal(t, 65, RunCalculationFunctionForMasses(sqrt, []int{4, 7}))
}
