package day5

// RunProgramCode runs the following program logic
// read the first 4 ints:
// 0 = Instructions
// 1 = position of first value
// 2 = position of seconds value
// 3 = position to write result to
//
// Instructions:
// 2 rightmost numbers == Opcode:
// 		1 - add first and second parameter
// 		2 - multiply first and second parameter
// 		3 - take input and write to position given in first parameter
// 		4 - output the first parameter
//		5 - if first param is not zero, set instruction pointer to value of second parameter
//		6 - if first param is zero, set instruction pointer to value of second parameter
//		7 - if first param < second param, store "1" to position of third params, else "0"
//		8 - if first param == second param, store "1" to position of third params, else "0"
// 		99 - exit
// The other value going RTL are the parameter modes for each parameter the opcode takes.
// 	0 - position mode (take value from this position)
// 	1 - value mode (directly take the value)
// If no parameter mode is given for a parameter, it is 0.
//
// After executing the opcode, step forward the amount of operations done and repeat.
func RunProgramCode(code []int, inputs *[]int, start int, output *[]int) []int {
	// default for input and output
	valuesInInstruction := 0
	instruction := code[start]
	opcode := instruction % 100

	getParameterValue := func(mode, position int) int {
		if mode == 1 {
			return code[position]
		}
		return code[code[position]]
	}

	switch opcode {
	case 1:
		valuesInInstruction = 4
		value1 := getParameterValue((instruction/100)%10, start+1)
		value2 := getParameterValue((instruction/100)/10, start+2)
		resultPosition := code[start+3]
		code[resultPosition] = value1 + value2
	case 2:
		valuesInInstruction = 4
		value1 := getParameterValue((instruction/100)%10, start+1)
		value2 := getParameterValue((instruction/100)/10, start+2)
		resultPosition := code[start+3]
		code[resultPosition] = value1 * value2
	case 3:
		valuesInInstruction = 2
		code[code[start+1]] = (*inputs)[0]
		if len(*inputs) > 1 {
			*inputs = append((*inputs)[:0], (*inputs)[1:]...)
		}
	case 4:
		valuesInInstruction = 2
		value1 := getParameterValue((instruction/100)%10, start+1)
		*output = append(*output, value1)
	case 5:
		value1 := getParameterValue((instruction/100)%10, start+1)
		if value1 != 0 {
			start = getParameterValue((instruction/100)/10, start+2)
		} else {
			valuesInInstruction = 3
		}
	case 6:
		value1 := getParameterValue((instruction/100)%10, start+1)
		if value1 == 0 {
			start = getParameterValue((instruction/100)/10, start+2)
		} else {
			valuesInInstruction = 3
		}
	case 7:
		valuesInInstruction = 4
		value1 := getParameterValue((instruction/100)%10, start+1)
		value2 := getParameterValue((instruction/100)/10, start+2)
		if value1 < value2 {
			code[code[start+3]] = 1
		} else {
			code[code[start+3]] = 0
		}
	case 8:
		valuesInInstruction = 4
		value1 := getParameterValue((instruction/100)%10, start+1)
		value2 := getParameterValue((instruction/100)/10, start+2)
		if value1 == value2 {
			code[code[start+3]] = 1
		} else {
			code[code[start+3]] = 0
		}
	case 99:
		return code
	default:
		return code
	}

	return RunProgramCode(code, inputs, start+valuesInInstruction, output)
}
