package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day1-the-tyranny-of-rocket-equation/day1"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	fmt.Println("===============================")
	masses, err := day1.LoadModulesMassesFromFile()
	if err != nil {
		fmt.Println("Unable to load spacecraft modules!")
		return
	}
	fuelPart1 := day1.RunCalculationFunctionForMasses(day1.CalculateRequiredFuel, masses)
	fmt.Println("The required fuel for all modules is:", fuelPart1)

	fmt.Println("")
	fmt.Println("Advent of Code - Day 1 - Part 2")
	fmt.Println("===============================")
	fuelPart2 := day1.RunCalculationFunctionForMasses(day1.CalculateRequiredFuelWithRecursion, masses)
	fmt.Println("The required fuel for all modules, including the mass of the fuel itself, is", fuelPart2)
}
