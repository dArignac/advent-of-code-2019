package part1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNearestCrossingDistance(t *testing.T) {
	input := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	assert.Equal(t, CalculateNearestCrossingDistance(input), 159)

	input = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	assert.Equal(t, CalculateNearestCrossingDistance(input), 137)
}

func TestSplitInputToPath(t *testing.T) {
	r1, e1 := splitInputToPath("R1,L2,X99")
	assert.Nil(t, e1)
	assert.Equal(t, r1, []path{path{"R", 1}, path{"L", 2}, path{"X", 99}})
}
