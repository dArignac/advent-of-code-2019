package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCalculationFunctionForMasses(t *testing.T) {
	sqrt := func(x int) int {
		return x * x
	}
	assert.Equal(t, 65, runCalculationFunctionForMasses(sqrt, []int{4, 7}))
}
