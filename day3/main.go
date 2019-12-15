package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dArignac/advent-of-code-2019/day3/part1"
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

	distance := part1.CalculateNearestCrossingDistance(input)
	if distance == -1 || distance == 0 {
		fmt.Println("Calculation failed")
	}
	fmt.Println("The shortest distance is", distance)
}
