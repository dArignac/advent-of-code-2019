package part2

import "github.com/dArignac/advent-of-code-2019/day1/part1"

func CalculateRequiredFuel(mass int) int {
	fuelSum := part1.CalculateRequiredFuel(mass)
	if fuelSum <= 0 {
		return fuelSum
	}
	return fuelSum + CalculateRequiredFuel(fuelSum)
}

func CalculateRequiredFuelForMultipleMasses(masses []int) int {
	result := 0
	for _, mass := range masses {
		result += CalculateRequiredFuel(mass)
	}
	return result
}
