package main

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day6-universal-orbit-map/day6"
	"github.com/dArignac/advent-of-code-2019/helper"
)

func main() {
	fmt.Println("Advent of Code - Day 6 - Part 1")
	fmt.Println("===============================")

	input, err := helper.LoadFileContent()
	if err != nil {
		fmt.Println("Unable to load input from file")
		return
	}

	tree := day6.CreateTree(input)
	counter := 0
	day6.RecursiveCountOrbits(&tree, &counter)
	fmt.Println("The total number of direct and indirect orbits is", counter)

	// fmt.Println("")
	// fmt.Println("Advent of Code - Day 6 - Part 2")
	// fmt.Println("===============================")
}
