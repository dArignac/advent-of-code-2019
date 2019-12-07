package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day1/part2"

	"github.com/dArignac/advent-of-code-2019/day1/part1"
)

func main() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	fmt.Println("===============================")
	spacecraftModules, err := part1.LoadSpacecraftModulesFromFile()
	if err != nil {
		fmt.Println("Unable to load spacecraft modules!")
		return
	}
	fuelPart1 := part1.CalculateRequiredFuelForMultipleModules(spacecraftModules)
	fmt.Println("The required fuel for all modules is:", fuelPart1)

	fmt.Println("")
	fmt.Println("Advent of Code - Day 1 - Part 2")
	fmt.Println("===============================")
	// FIXME use the loaded masses
	fuelPart2 := part2.CalculateRequiredFuelWithRecursion(1969)
	fmt.Println("The required fuel for all modules, including the mass of the fuel itself, is", fuelPart2)
}
