package main

import "fmt"

import "github.com/dArignac/advent-of-code-2019/day2/part1"

func main() {
	fmt.Println("Advent of Code - Day 2 - Part 1")
	fmt.Println("===============================")
	// load the inputs as string
	codeString, err := part1.LoadProgramCodeFromFile()
	if err != nil {
		fmt.Println("Unable to load program code from file")
		return
	}

	// convert inputs to []int slice
	code, err := part1.ConvertProgramCode(codeString)
	if err != nil {
		fmt.Println("Unable to convert program code")
		return
	}

	// do the required replacements
	code[1] = 12
	code[2] = 2

	// fix the code
	code = part1.FixIntcodeProgramCode(code)

	// FIXME remove
	fmt.Println(code)

	fmt.Println("Value on position 0 is", code[0])
}
