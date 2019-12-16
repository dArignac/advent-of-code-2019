package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNearestCrossingWithDistance(t *testing.T) {
	w0, err0 := WireInstructionsToCoordinates("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w1, err1 := WireInstructionsToCoordinates("U62,R66,U55,R34,D71,R55,D58,R83")
	assert.Nil(t, err0)
	assert.Nil(t, err1)

	r1 := GetNearestCrossingWithDistance(GetCrossings(w0, w1))
	assert.Equal(t, 159.0, r1)

	w0, err0 = WireInstructionsToCoordinates("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w1, err1 = WireInstructionsToCoordinates("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	assert.Nil(t, err0)
	assert.Nil(t, err1)

	r2 := GetNearestCrossingWithDistance(GetCrossings(w0, w1))
	assert.Equal(t, 135.0, r2)
}

func TestCalculateShortestStepCrossing(t *testing.T) {
	// FIXME impl
}

func TestSplitInputToPath(t *testing.T) {
	r1, e1 := splitInputToPath("R1,L2,X99")
	assert.Nil(t, e1)
	assert.Equal(t, r1, []path{path{"R", 1}, path{"L", 2}, path{"X", 99}})
}

func TestPathToPoint(t *testing.T) {
	s1 := Point{x: 0, y: 0}
	p1 := path{direction: "R", length: 3}
	e1 := []Point{{x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}}
	r1 := pathToPoints(p1, s1)
	assert.Equal(t, e1, r1)

	s2 := Point{x: -66, y: 128}
	p2 := path{direction: "R", length: 3}
	e2 := []Point{{x: -65, y: 128}, {x: -64, y: 128}, {x: -63, y: 128}}
	r2 := pathToPoints(p2, s2)
	assert.Equal(t, e2, r2)

	s3 := Point{x: 0, y: 0}
	p3 := path{direction: "D", length: 2}
	e3 := []Point{{x: 0, y: -1}, {x: 0, y: -2}}
	r3 := pathToPoints(p3, s3)
	assert.Equal(t, e3, r3)

	s4 := Point{x: 7, y: 1}
	p4 := path{direction: "D", length: 3}
	e4 := []Point{{x: 7, y: 0}, {x: 7, y: -1}, {x: 7, y: -2}}
	r4 := pathToPoints(p4, s4)
	assert.Equal(t, e4, r4)

	s5 := Point{x: 0, y: 0}
	p5 := path{direction: "L", length: 2}
	e5 := []Point{{x: -1, y: 0}, {x: -2, y: 0}}
	r5 := pathToPoints(p5, s5)
	assert.Equal(t, e5, r5)

	s6 := Point{x: -7, y: 6}
	p6 := path{direction: "L", length: 3}
	e6 := []Point{{x: -8, y: 6}, {x: -9, y: 6}, {x: -10, y: 6}}
	r6 := pathToPoints(p6, s6)
	assert.Equal(t, e6, r6)

	s7 := Point{x: 0, y: 0}
	p7 := path{direction: "U", length: 2}
	e7 := []Point{{x: 0, y: 1}, {x: 0, y: 2}}
	r7 := pathToPoints(p7, s7)
	assert.Equal(t, e7, r7)

	s8 := Point{x: -7, y: -2}
	p8 := path{direction: "U", length: 3}
	e8 := []Point{{x: -7, y: -1}, {x: -7, y: 0}, {x: -7, y: 1}}
	r8 := pathToPoints(p8, s8)
	assert.Equal(t, e8, r8)
}

func TestContainsElement(t *testing.T) {
	h1 := []Point{{x: 3, y: 4}, {x: 5, y: 4}, {x: 77, y: -3}}
	assert.True(t, containsElement(h1, Point{x: 5, y: 4}))

	h2 := []Point{{x: 3, y: 4}, {x: 5, y: 4}, {x: 77, y: -3}}
	assert.False(t, containsElement(h2, Point{x: 7, y: -4}))

	h3 := []Point{{x: -7, y: 9}, {x: -166, y: 18}, {x: -77, y: -3}}
	assert.True(t, containsElement(h3, Point{x: -77, y: -3}))
}

func TestInsertPathCoordinates(t *testing.T) {
	p1 := path{direction: "U", length: 2}
	p2 := []Point{{x: 7, y: 3}, {x: 8, y: 3}}
	e1 := []Point{{x: 7, y: 3}, {x: 8, y: 3}, {x: 8, y: 4}, {x: 8, y: 5}}
	assert.Equal(t, e1, insertPathCoordinates(p1, p2))

	// the coordinates need at least one element
	assert.Panics(t, func() { insertPathCoordinates(p1, []Point{}) })
}

func TestWireInstructionsToCoordinates(t *testing.T) {
	result1, error1 := WireInstructionsToCoordinates("U1,R1,D1,L1")
	expected1 := []Point{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 1}, {x: 1, y: 0}, {x: 0, y: 0}}
	assert.Nil(t, error1)
	assert.Equal(t, expected1, result1)
}

func TestGetCrossings(t *testing.T) {
	w1 := []Point{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 2}, {2, 3}}
	w2 := []Point{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}, {3, 2}}

	c1 := GetCrossings(w1, w2)
	e1 := []CrossingPoint{{point: Point{1, 1}, steps: 0}, {point: Point{2, 2}, steps: 0}}

	assert.Equal(t, e1, c1)
}
