package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day1-the-tyranny-of-rocket-equation/day1"
	"github.com/dArignac/advent-of-code-2019/helper"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	fmt.Println("===============================")

	content, err := helper.LoadFileContent()
	if err != nil {
		fmt.Println("Unable to load input file!")
		return
	}

	masses, err := helper.ConvertStringArrayToIntArray(content)
	if err != nil {
		fmt.Println("Unable to convert input file!")
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
