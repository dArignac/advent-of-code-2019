package helper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// LoadFileContent loads the contents of the input file as string
func LoadFileContent() (string, error) {
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

// ConvertProgramCode converts the given program code into an integer slice
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