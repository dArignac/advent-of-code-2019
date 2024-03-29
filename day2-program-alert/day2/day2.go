package day2

import (
	"fmt"
	"github.com/dArignac/advent-of-code-2019/helper"
)

// FixIntcodeProgramCode fixes the program code by the following logic
// read the first 4 ints:
// 0 = Opcode
// 1 = position of first value
// 2 = position of seconds value
// 3 = position to write result to
//
// Opcodes:
// 1 - add first and second value
// 2 - multiply first and second value
// 99 - exit
//
// After executing the opcode, step forward 4 positions and repeat.
func FixIntcodeProgramCode(code []int, start int) []int {
	opcode := code[start]

	if opcode == 99 {
		return code
	}

	value1 := code[code[start+1]]
	value2 := code[code[start+2]]
	targetPosition := code[start+3]
	nextStart := start + 4
	length := len(code)

	switch opcode {
	case 1:
		code[targetPosition] = value1 + value2
	case 2:
		code[targetPosition] = value1 * value2
	}

	if nextStart < length {
		code = FixIntcodeProgramCode(code, nextStart)
	}
	return code
}

// FindInputsForDesiredOutput searches the input values to get the given output.
func FindInputsForDesiredOutput(intCode string, output int) (int, int) {
	// grab the initial Intcode
	// convert inputs to []int slice
	code, err := helper.SplitStringToIntArray(intCode)
	if err != nil {
		fmt.Println("Unable to convert program code")
		return 0, 0
	}

	// value range for noun and verb is 0-99
	noun := 0
	verb := 0
	for noun < 100 {
		for verb < 100 {
			iterationCode := make([]int, len(code))
			copy(iterationCode, code)
			// iterationCode := append([]int(nil), code...)
			iterationCode[1] = noun
			iterationCode[2] = verb

			result := FixIntcodeProgramCode(iterationCode, 0)
			if result[0] == output {
				return noun, verb
			}

			verb++
		}
		verb = 0
		noun++
	}

	// FIXME remove
	return 0, 0
}
