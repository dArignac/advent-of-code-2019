package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func CalculateRequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

func CalculateRequiredFuelForMultipleModules(modules []int) int {
	result := 0
	for _, v := range modules {
		result += CalculateRequiredFuel(v)
	}
	return result
}

func LoadSpacecraftModulesFromFile() ([]int, error) {
	file, err := os.Open("spacecraft-modules.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var spacecraftModules []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		module, _ := strconv.Atoi(scanner.Text())
		spacecraftModules = append(spacecraftModules, module)
	}

	return spacecraftModules, nil
}

func main() {
	fmt.Println("Advent of Code - Day 1")
	spacecraftModules, err := LoadSpacecraftModulesFromFile()
	if err != nil {
		fmt.Println("Unable to load spacecraft modules!")
		return
	}
	fmt.Println("The required fuel for all modules is:", CalculateRequiredFuelForMultipleModules(spacecraftModules))
}
