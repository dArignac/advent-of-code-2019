package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dArignac/advent-of-code-2019/day3/part1"
	"github.com/dArignac/advent-of-code-2019/day3/part2"
)

func loadInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content, nil
}

func main() {
	fmt.Println("Advent of Code - Day 3 - Part 1")
	fmt.Println("===============================")

	input, err := loadInput()
	if err != nil {
		fmt.Println("Error occurred", err)
		return
	}

	// transform wire instructions to coordinate
	coordinatesWire0, err := part1.WireInstructionsToCoordinates(input[0])
	if err != nil {
		fmt.Println("Tranforming instructions for wire 1 to coordinates failed")
		return
	}
	coordinatesWire1, err := part1.WireInstructionsToCoordinates(input[1])
	if err != nil {
		fmt.Println("Tranforming instructions for wire 2 to coordinates failed")
		return
	}

	// calculate the nearest crossing distance
	distance := part1.CalculateNearestCrossingDistance(coordinatesWire0, coordinatesWire1)
	if distance == -1 || distance == 0 {
		fmt.Println("Calculation failed, distance was", distance)
	}
	fmt.Println("The shortest distance is", distance)

	fmt.Println("")
	fmt.Println("Advent of Code - Day 3 - Part 2")
	fmt.Println("===============================")

	steps := part2.CalculateShortestStepCrossing()
	fmt.Println("Shorted steps required are", steps, "steps.")
}
