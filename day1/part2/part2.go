package part2

import "github.com/dArignac/advent-of-code-2019/day1/part1"

func CalculateRequiredFuelWithRecursion(mass int) int {
	fuelSum := part1.CalculateRequiredFuel(mass)
	if fuelSum <= 0 {
		return fuelSum
	}
	return fuelSum + CalculateRequiredFuelWithRecursion(fuelSum)
}
