package part1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNearestCrossingDistance(t *testing.T) {
	input := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	d1, _ := CalculateNearestCrossingDistance(input)
	assert.Equal(t, d1, 159)

	input = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	d2, _ := CalculateNearestCrossingDistance(input)
	assert.Equal(t, d2, 135)
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

func TestGetNearestCrossingWithDistance(t *testing.T) {
	c1 := []Point{{x: -100, y: 2}, {x: 100, y: 1}}
	p1, d1 := getNearestCrossingWithDistance(c1)
	assert.Equal(t, Point{x: 100, y: 1}, p1)
	assert.Equal(t, 101.0, d1)

	c2 := []Point{{x: -5, y: -30}, {x: 7, y: -11}, {x: 17, y: -55}}
	p2, d2 := getNearestCrossingWithDistance(c2)
	assert.Equal(t, Point{x: 7, y: -11}, p2)
	assert.Equal(t, 18.0, d2)
}
