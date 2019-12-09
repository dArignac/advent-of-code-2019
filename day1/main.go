package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day1/part2"

	"github.com/dArignac/advent-of-code-2019/day1/part1"
)

func runCalculationFunctionForMasses(calculation func(int) int, masses []int) int {
	result := 0
	for _, mass := range masses {
		result += calculation(mass)
	}
	return result
}

func main() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	fmt.Println("===============================")
	masses, err := part1.LoadModulesMassesFromFile()
	if err != nil {
		fmt.Println("Unable to load spacecraft modules!")
		return
	}
	fuelPart1 := runCalculationFunctionForMasses(part1.CalculateRequiredFuel, masses)
	fmt.Println("The required fuel for all modules is:", fuelPart1)

	fmt.Println("")
	fmt.Println("Advent of Code - Day 1 - Part 2")
	fmt.Println("===============================")
	fuelPart2 := runCalculationFunctionForMasses(part2.CalculateRequiredFuel, masses)
	fmt.Println("The required fuel for all modules, including the mass of the fuel itself, is", fuelPart2)
}
