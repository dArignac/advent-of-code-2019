package part2

import (
	"fmt"

	"github.com/dArignac/advent-of-code-2019/day2/part1"
)

func FindInputsForDesiredOutput(intCode string, output int) (int, int) {
	// grab the inital Intcode
	// convert inputs to []int slice
	code, err := part1.ConvertProgramCode(intCode)
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

			result := part1.FixIntcodeProgramCode(iterationCode, 0)
			if result[0] == output {
				return noun, verb
			}

			verb += 1
		}
		verb = 0
		noun += 1
	}

	// FIXME remove
	return 0, 0
}
