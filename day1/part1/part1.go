package part1

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func CalculateRequiredFuel(mass int) int {
	fuel := int(math.Floor(float64(mass/3))) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func CalculateRequiredFuelForMultipleMasses(masses []int) int {
	result := 0
	for _, mass := range masses {
		result += CalculateRequiredFuel(mass)
	}
	return result
}

func LoadModulesMassesFromFile() ([]int, error) {
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
