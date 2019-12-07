package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuel(t *testing.T) {
	assert.Equal(t, 2, CalculateRequiredFuel(12))
	assert.Equal(t, 2, CalculateRequiredFuel(14))
	assert.Equal(t, 654, CalculateRequiredFuel(1969))
	assert.Equal(t, 33583, CalculateRequiredFuel(100756))
}
