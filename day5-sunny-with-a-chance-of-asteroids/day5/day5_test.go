package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixIntcodeProgramCode(t *testing.T) {
	code := []int{99}
	output := []int{7}
	assert.Equal(t, []int{99}, RunProgramCode(code, 0, 0, &output))
	assert.Equal(t, []int{7}, output)

	output = []int{}

	// all params as positions
	code = []int{1, 0, 0, 0, 99}
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

	code = []int{2, 3, 0, 3, 99}
	assert.Equal(t, []int{2, 3, 0, 6, 99}, RunProgramCode(code, 0, 0, &output))

	code = []int{2, 4, 4, 5, 99, 0}
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, RunProgramCode(code, 0, 0, &output))

	code = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, RunProgramCode(code, 0, 0, &output))

	// code = []int{3, 0, 4, 0, 99}
	// RunProgramCode(code, 666, 0, &output)
	// assert.Equal(t, output, []int{666})

	// code = []int{1002, 4, 3, 4, 33}
	// output = []int{}
	// assert.Equal(t, []int{1002, 4, 3, 4, 99}, RunProgramCode(code, 0, 0, &output))

	// code = []int{1101, 100, -1, 4, 0}
	// assert.Equal(t, []int{1101, 100, -1, 4, 99}, RunProgramCode(code, 0, 0, &output))
}
