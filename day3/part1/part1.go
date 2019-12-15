package part1

import (
	"strconv"
	"strings"
)

// find crossings
// calculate path to crossings
// choose the one with the shortest path

type path struct {
	direction string
	length    int
}

// CalculateNearestCrossingDistance calculates the distance to the shortest crossing
func CalculateNearestCrossingDistance(lines []string) int {
	return 0
}

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

// TODO split string to elements (use a struct?)
// TODO transfer to a matrix (use array or struct)
// TODO on insertion check for crossings and remember
// TODO calculate distances for crossings and return the shortest
