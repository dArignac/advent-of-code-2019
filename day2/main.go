package main

import "fmt"

import "github.com/dArignac/advent-of-code-2019/day2/part1"

import "github.com/dArignac/advent-of-code-2019/day2/part2"

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
	code = part1.FixIntcodeProgramCode(code, 0)

	// ...and the result is:
	fmt.Println("Value on position 0 is", code[0])

	fmt.Println("")
	fmt.Println("Advent of Code - Day 2 - Part 2")
	fmt.Println("===============================")

	noun, verb := part2.FindInputsForDesiredOutput(codeString, 19690720)
	fmt.Println("100 * noun * verb for output 19690720")
	fmt.Printf("Noun: %d, Verb: %d, Result: %d", noun, verb, 100*noun+verb)
}
