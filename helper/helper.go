package helper

import (
	"bufio"
	"os"
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
