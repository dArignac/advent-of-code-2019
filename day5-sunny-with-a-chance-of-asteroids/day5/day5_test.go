package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunProgramCodeNoop(t *testing.T) {
	code := []int{99}
	output := []int{7}
	assert.Equal(t, []int{99}, RunProgramCode(code, 0, 0, &output))
	assert.Equal(t, []int{7}, output)
}

func TestRunProgramCodeAddition(t *testing.T) {
	output := []int{}

	// all params as positions
	code := []int{1, 0, 0, 0, 99}
	assert.Equal(t, []int{2, 0, 0, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// all params as values
	code = []int{1101, 3, 4, 0, 99}
	assert.Equal(t, []int{7, 3, 4, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// mixed scenario 1
	code = []int{1001, 4, 9, 0, 99}
	assert.Equal(t, []int{99 + 9, 4, 9, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// mixed scenario 2
	code = []int{101, 3, 0, 0, 99}
	assert.Equal(t, []int{3 + 101, 3, 0, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// day 3
	code = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// day 5
	code = []int{1101, 100, -1, 4, 0}
	assert.Equal(t, []int{1101, 100, -1, 4, 99}, RunProgramCode(code, 0, 0, &output))
}

func TestRunProgramCodeMultiplication(t *testing.T) {
	output := []int{}
	code := []int{2, 3, 0, 3, 99}
	assert.Equal(t, []int{2, 3, 0, 6, 99}, RunProgramCode(code, 0, 0, &output))

	code = []int{2, 4, 4, 5, 99, 0}
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, RunProgramCode(code, 0, 0, &output))

	code = []int{1002, 4, 3, 4, 33}
	assert.Equal(t, []int{1002, 4, 3, 4, 99}, RunProgramCode(code, 0, 0, &output))
}

func TestRunProgramCodeInputOutput(t *testing.T) {
	output := []int{}
	code := []int{3, 0, 4, 0, 99}
	assert.Equal(t, code, RunProgramCode(code, 666, 0, &output))
	assert.Equal(t, []int{666}, output)
}

func TestRunProgramCodeCompare8(t *testing.T) {
	output := []int{}
	code := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	RunProgramCode(code, 8, 0, &output)
	assert.Equal(t, []int{1}, output)

	output = []int{}
	RunProgramCode(code, 7, 0, &output)
	assert.Equal(t, []int{0}, output)

	output = []int{}
	code = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	RunProgramCode(code, 7, 0, &output)
	assert.Equal(t, []int{1}, output)

	output = []int{}
	RunProgramCode(code, 9, 0, &output)
	assert.Equal(t, []int{0}, output)

	output = []int{}
	code = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	RunProgramCode(code, 8, 0, &output)
	assert.Equal(t, []int{1}, output)

	output = []int{}
	RunProgramCode(code, 7, 0, &output)
	assert.Equal(t, []int{0}, output)

	output = []int{}
	code = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	RunProgramCode(code, 7, 0, &output)
	assert.Equal(t, []int{1}, output)

	output = []int{}
	RunProgramCode(code, 9, 0, &output)
	assert.Equal(t, []int{0}, output)
}

func TestRunProgramCodeJumps(t *testing.T) {
	output := []int{}
	code := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	RunProgramCode(code, 0, 0, &output)
	assert.Equal(t, []int{0}, output)

	output = []int{}
	code = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	RunProgramCode(code, 1, 0, &output)
	assert.Equal(t, []int{1}, output)

	output = []int{}
	code = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	RunProgramCode(code, 0, 0, &output)
	assert.Equal(t, []int{0}, output)

	output = []int{}
	code = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	RunProgramCode(code, 1, 0, &output)
	assert.Equal(t, []int{1}, output)
}

func TestRunProgramCodeLarge(t *testing.T) {
	// output := []int{}
	// code := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	// RunProgramCode(code, 7, 0, &output)
	// assert.Equal(t, []int{999}, output)

	// 	output = []int{}
	// 	code = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	// 	RunProgramCode(code, 8, 0, &output)
	// 	assert.Equal(t, []int{1000}, output)

	// 	output = []int{}
	// 	code = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	// 	RunProgramCode(code, 9, 0, &output)
	// 	assert.Equal(t, []int{1001}, output)
}
