package main

import "fmt"

func main() {
	fmt.Println("Advent of Code - Day 1 - Part 1")
	spacecraftModules, err := loadSpacecraftModulesFromFile()
	if err != nil {
		fmt.Println("Unable to load spacecraft modules!")
		return
	}
	fmt.Println("The required fuel for all modules is:", calculateRequiredFuelForMultipleModules(spacecraftModules))
}
