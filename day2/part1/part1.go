package part1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadProgramCodeFromFile() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = scanner.Text()
	}

	return content, nil
}

func ConvertProgramCode(code string) ([]int, error) {
	var result []int

	for _, v := range strings.Split(code, ",") {
		step, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, step)
	}

	return result, nil
}

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
func FixIntcodeProgramCode(code []int) []int {
	return code
}
