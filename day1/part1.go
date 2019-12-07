package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func calculateRequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

func calculateRequiredFuelForMultipleModules(modules []int) int {
	result := 0
	for _, v := range modules {
		result += calculateRequiredFuel(v)
	}
	return result
}

func loadSpacecraftModulesFromFile() ([]int, error) {
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
