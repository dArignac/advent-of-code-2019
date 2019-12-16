package part1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type path struct {
	direction string
	length    int
}

// Point reflects a position within the wire coordindate system
type Point struct {
	x int16
	y int16
}

// CalculateNearestCrossingDistance calculates the distance to the shortest crossing
func CalculateNearestCrossingDistance(coordinatesWire0 []Point, coordinatesWire1 []Point) int {
	fmt.Println("Wire one contains", len(coordinatesWire0), "coordinates")
	fmt.Println("Wire two contains", len(coordinatesWire1), "coordinates")

	// calculate crossings by checking if wire 2 has any points that also occure in wire 1
	var crossings []Point
	for _, point := range coordinatesWire1 {
		if containsElement(coordinatesWire0, point) {
			crossings = append(crossings, point)
		}
	}
	// remove first element as this is the starting point 0,0
	crossings = crossings[1:]

	fmt.Println("Found", len(crossings), "crossings between the 2 wires")

	// calculate taxi distance for existing crossing, return the shortest
	_, distance := getNearestCrossingWithDistance(crossings)

	return int(distance)
}

// WireInstructionsToCoordinates takes stringified instructions for a wire and transforms them to coordinates reflected by the Point struct.
func WireInstructionsToCoordinates(instructions string) ([]Point, error) {
	coordinates := []Point{{x: 0, y: 0}}

	paths, err := splitInputToPath(instructions)
	if err != nil {
		return nil, err
	}

	for _, path := range paths {
		coordinates = insertPathCoordinates(path, coordinates)
	}

	return coordinates, nil
}

func getNearestCrossingWithDistance(crossings []Point) (Point, float64) {
	var nearestPoint Point
	nearestDistance := -1.0

	for _, point := range crossings {
		distance := math.Abs(float64(point.x)) + math.Abs(float64(point.y))
		if nearestDistance == -1.0 {
			nearestDistance = distance
		} else {
			if distance < nearestDistance {
				nearestPoint = point
				nearestDistance = distance
			}
		}
	}
	return nearestPoint, nearestDistance
}

func containsElement(haystack []Point, needle Point) bool {
	for _, element := range haystack {
		if element.x == needle.x && element.y == needle.y {
			return true
		}
	}
	return false
}

// insertPathCoordinates creates points from the paths and inserts them into coordinates
func insertPathCoordinates(path path, coordinates []Point) []Point {
	// create points slice from path
	points := pathToPoints(path, coordinates[len(coordinates)-1])

	for _, point := range points {
		coordinates = append(coordinates, point)
	}

	return coordinates
}

func pathToPoints(path path, start Point) []Point {
	result := make([]Point, 0)
	for i := 1; i <= path.length; i++ {
		switch path.direction {
		case "R":
			result = append(result, Point{x: start.x + int16(i), y: start.y})
		case "D":
			result = append(result, Point{x: start.x, y: start.y - int16(i)})
		case "L":
			result = append(result, Point{x: start.x - int16(i), y: start.y})
		case "U":
			result = append(result, Point{x: start.x, y: start.y + int16(i)})
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
