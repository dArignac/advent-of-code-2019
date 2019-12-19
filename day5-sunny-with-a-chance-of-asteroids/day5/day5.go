package day5

import "fmt"

// RunProgramCode runs the following program logic
// read the first 4 ints:
// 0 = Instructions
// 1 = position of first value
// 2 = position of seconds value
// 3 = position to write result to
//
// Instructions:
// 2 rightmost numbers == Opcode:
// 		1 - add first and second value
// 		2 - multiply first and second value
// 		3 - take input and write position given on next param
// 		4 - output the next param
// 		99 - exit
// The other value going RTL are the parameter modes for each parameter the opcode takes.
// 	0 - position mode (take value from this position)
// 	1 - value mode (directly take the value)
// If no parameter mode is given for a parameter, it is 0.
//
// After executing the opcode, step forward the amount of operations done and repeat.
func RunProgramCode(code []int, input int, start int, output *[]int) []int {
	// default for add and multiply
	valuesInInstruction := 4
	instruction := code[start]
	opcode := instruction % 100

	getParameterValue := func(mode, position int) int {
		fmt.Println("mode", mode, "position", position)
		if mode == 1 {
			return code[position]
		}
		return code[code[position]]
	}

	if opcode == 99 {
		return code
	}

	// addition and multiplication
	if opcode == 1 || opcode == 2 {
		value1 := getParameterValue((instruction/100)%10, start+1)
		value2 := getParameterValue((instruction/100)/10, start+2)
		resultPosition := code[start+3]
		if opcode == 1 {
			code[resultPosition] = value1 + value2
		} else {
			code[resultPosition] = value1 * value2
		}
	}

	if opcode == 2 || opcode == 3 {
		// FIXME implement
	}

	// invalid, just return
	if opcode != 99 && (opcode > 3 || opcode < 0) {
		return code
	}

	return RunProgramCode(code, input, start+valuesInInstruction, output)
}
