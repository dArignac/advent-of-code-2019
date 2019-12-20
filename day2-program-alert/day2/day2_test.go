package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixIntcodeProgramCode(t *testing.T) {
	assert.Equal(t, []int{2, 0, 0, 0, 99}, FixIntcodeProgramCode([]int{1, 0, 0, 0, 99}, 0))
	assert.Equal(t, []int{2, 3, 0, 6, 99}, FixIntcodeProgramCode([]int{2, 3, 0, 3, 99}, 0))
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, FixIntcodeProgramCode([]int{2, 4, 4, 5, 99, 0}, 0))
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, FixIntcodeProgramCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0))
}
