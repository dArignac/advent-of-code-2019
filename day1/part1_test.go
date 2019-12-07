package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuel(t *testing.T) {
	assert.Equal(t, 2, calculateRequiredFuel(12))
	assert.Equal(t, 2, calculateRequiredFuel(14))
	assert.Equal(t, 654, calculateRequiredFuel(1969))
	assert.Equal(t, 33583, calculateRequiredFuel(100756))
}

func TestCalculateRequiredFuelForMultipleModules(t *testing.T) {
	modules := []int{12, 14, 1969, 100756}
	assert.Equal(t, 2+2+654+33583, calculateRequiredFuelForMultipleModules(modules))
}
