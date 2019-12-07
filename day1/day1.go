package main

import (
	"fmt"
	"math"
)

func CalculateRequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

func main() {
	fmt.Println("Advent of Code - Day 1")
}
