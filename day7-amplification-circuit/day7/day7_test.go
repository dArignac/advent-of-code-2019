package day7

import "testing"

import "github.com/stretchr/testify/assert"

func TestFindLargestOutputSignalForThrusters43210(t *testing.T) {
	code := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	result := FindLargestOutputSignalForThrusters(code)
	assert.Equal(t, Phases{4, 3, 2, 1, 0, 43210}, result)
}

func TestFindLargestOutputSignalForThrusters54321(t *testing.T) {
	code := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	result := FindLargestOutputSignalForThrusters(code)
	assert.Equal(t, Phases{0, 1, 2, 3, 4, 54321}, result)
}

func TestFindLargestOutputSignalForThrusters65210(t *testing.T) {
	code := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	result := FindLargestOutputSignalForThrusters(code)
	assert.Equal(t, Phases{1, 0, 4, 3, 2, 65210}, result)
}

func TestGeneratePermutations(t *testing.T) {
	result := generatePermutations([]int{0, 1})
	expected := [][]int{
		[]int{0, 1},
		[]int{1, 0},
	}
	assert.Equal(t, expected, result)

	result = generatePermutations([]int{0, 1, 2})
	expected = [][]int{
		[]int{0, 1, 2},
		[]int{1, 0, 2},
		[]int{2, 1, 0},
		[]int{1, 2, 0},
		[]int{2, 0, 1},
		[]int{0, 2, 1},
	}
	assert.Equal(t, expected, result)
}
