package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day7-amplification-circuit/day7"
	"github.com/dArignac/advent-of-code-2019/helper"
)

func main() {
	fmt.Println("Advent of Code - Day 7 - Part 1")
	fmt.Println("===============================")

	input, err := helper.LoadFileContent()
	if err != nil {
		fmt.Println("Unable to load input from file")
		return
	}

	result := day7.FindLargestOutputSignalForThrusters(input[0])
	fmt.Println("Result is", result.ToString())
}
