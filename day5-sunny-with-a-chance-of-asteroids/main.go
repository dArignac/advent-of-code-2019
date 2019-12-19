package main

import (
	"fmt"
	"github.com/dArignac/advent-of-code-2019/day5-sunny-with-a-chance-of-asteroids/day5"
	"github.com/dArignac/advent-of-code-2019/helper"
)

func main() {
	fmt.Println("Advent of Code - Day 5 - Part 1")
	fmt.Println("===============================")

	input, err := helper.LoadFileContent()
	if err != nil {
		fmt.Println("Unable to load input from file")
		return
	}

	code, err := helper.ConvertProgramCode(input)
	if err != nil {
		fmt.Println("Unable to convert input to program code")
		return
	}

	fmt.Println("Now running the program code with initial position 1:")
	output := []int{}
	day5.RunProgramCode(code, 1, 0, &output)
	fmt.Println("Outputs are")
	fmt.Println(output)

	// fmt.Println("")
	// fmt.Println("Advent of Code - Day 5 - Part 2")
	// fmt.Println("===============================")
}
