package part1

import (
	"strconv"
	"strings"
)

type path struct {
	direction string
	length    int
}

type point struct {
	x int16
	y int16
}

// CalculateNearestCrossingDistance calculates the distance to the shortest crossing
func CalculateNearestCrossingDistance(wires []string) int {
	coordinates := []point{{x: 0, y: 0}}
	var crossings []point

	for _, wire := range wires {
		paths, err := splitInputToPath(wire)
		if err != nil {
			return -1
		}

		for _, path := range paths {
			insertPathCoordinates(path, &coordinates, &crossings)
		}
	}

	// FIXME calculate taxi distance for existing crossing, return the shortest

	return 0
}

func insertPathCoordinates(path path, coordinates *[]point, crossings *[]point) {
	// create points slice from path
	// check if any of the points exists, that then marks a crossing, add this to crossing
	// add the points to the coordinates
}

func pathToPoints(path path, start point) []point {
	result := make([]point, 0)
	for i := 1; i <= path.length; i++ {
		switch path.direction {
		case "R":
			result = append(result, point{x: start.x + int16(i), y: start.y})
		case "D":
			result = append(result, point{x: start.x, y: start.y - int16(i)})
		case "L":
			result = append(result, point{x: start.x - int16(i), y: start.y})
		case "U":
			result = append(result, point{x: start.x, y: start.y + int16(i)})
		}
	}
	return result
}

// splitInputToPath splits the given string into "path" instances
func splitInputToPath(input string) ([]path, error) {
	var paths []path
	elements := strings.Split(input, ",")
	for _, v := range elements {
		length, err := strconv.Atoi(v[1:])
		if err != nil {
			return nil, err
		}
		paths = append(paths, path{v[:1], length})
	}
	return paths, nil
}
