package day1

import (
	"math"
)

// CalculateRequiredFuel returns the required fuel for the given mass.
func CalculateRequiredFuel(mass int) int {
	fuel := int(math.Floor(float64(mass/3))) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

// CalculateRequiredFuelWithRecursion returns the required fuel for the given mass and takes the mass of the fuel itself into account.
func CalculateRequiredFuelWithRecursion(mass int) int {
	fuelSum := CalculateRequiredFuel(mass)
	if fuelSum <= 0 {
		return fuelSum
	}
	return fuelSum + CalculateRequiredFuelWithRecursion(fuelSum)
}

// RunCalculationFunctionForMasses runs the given function for each element of the given masses and sums the return values.
func RunCalculationFunctionForMasses(calculation func(int) int, masses []int) int {
	result := 0
	for _, mass := range masses {
		result += calculation(mass)
	}
	return result
}
