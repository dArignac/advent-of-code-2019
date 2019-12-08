package part1

import "testing"

import "github.com/stretchr/testify/assert"

func TestConvertProgramCode(t *testing.T) {
	r1, e1 := ConvertProgramCode("0,1,2,3")
	assert.Nil(t, e1)
	assert.Equal(t, r1, []int{0, 1, 2, 3})

	r2, e2 := ConvertProgramCode("7,129,18,44")
	assert.Nil(t, e2)
	assert.Equal(t, r2, []int{7, 129, 18, 44})

	r3, e3 := ConvertProgramCode("")
	assert.NotNil(t, e3)
	assert.Equal(t, []int(nil), r3)
}

func TestFixIntcodeProgramCode(t *testing.T) {
	assert.Equal(t, FixIntcodeProgramCode([]int{1, 0, 0, 0, 99}, 0), []int{2, 0, 0, 0, 99})
	assert.Equal(t, FixIntcodeProgramCode([]int{2, 3, 0, 3, 99}, 0), []int{2, 3, 0, 6, 99})
	assert.Equal(t, FixIntcodeProgramCode([]int{2, 4, 4, 5, 99, 0}, 0), []int{2, 4, 4, 5, 99, 9801})
	assert.Equal(t, FixIntcodeProgramCode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0), []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
