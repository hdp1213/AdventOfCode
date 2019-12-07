package day03

import (
	"bufio"
	"fmt"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

const (
	horizontal = iota
	vertical = iota
)

const (
	positive = iota
	negative = iota
)

type point struct {
	x int
	y int
}

func (point *point) manhattanDistance() int {
	return abs(point.x) + abs(point.y)
}

type pointArray []point

type lineSegment struct {
	start point
	end point
	orientation int
	direction int
}

func (points pointArray) contains(testPoint point) bool {
	for _, point := range points {
		if testPoint.x == point.x && testPoint.y == point.y {
			return true
		}
	}
	return false
}

func readWirePaths(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result, scanner.Err()
}

func getWirePathElements(wirePath string) ([]string, error) {
	stringReader := strings.NewReader(wirePath)
	scanner := bufio.NewScanner(stringReader)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)

	var pathElements []string
	for scanner.Scan() {
		elem := scanner.Text()
		pathElements = append(pathElements, elem)
	}

	return pathElements, scanner.Err()
}

func getCoords(wireElements []string) (points pointArray, err error) {
	// As a reminder, x is R/L, y is U/D. +ve x -> R, +ve y -> U
	newPoint := point {x:0,y:0}
	points = make(pointArray, len(wireElements) + 1)
	points[0] = newPoint

	for i, element := range wireElements {
		direction := element[0]
		lengthString := element[1:len(element)]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return points, err
		}

		if direction == 'R' {
			newPoint.x = points[i].x + length
		}
		if direction == 'L' {
			newPoint.x = points[i].x - length
		}

		if direction == 'U' {
			newPoint.y = points[i].y + length
		}
		if direction == 'D' {
			newPoint.y = points[i].y - length
		}

		points[i + 1] = newPoint
	}

	return
}

func getLineSegments(points pointArray) ([]lineSegment, error) {
	lineSegments := make([]lineSegment, len(points) - 1)

	for i := 1; i < len(points); i++ {
		lineSegment, err := newLineSegment(points[i - 1], points[i])

		if err != nil {
			return nil, err
		}

		lineSegments[i - 1] = lineSegment
	}

	return lineSegments, nil
}

func processWirePath(wirePath string) ([]lineSegment, error) {
	pathElements, err := getWirePathElements(wirePath)
	if err != nil {
		return nil, err
	}

	coords, err := getCoords(pathElements)
	if err != nil {
		return nil, err
	}

	lineSegments, err := getLineSegments(coords)
	if err != nil {
		return nil, err
	}

	return lineSegments, nil
}

func doesIntersect(seg1 lineSegment, seg2 lineSegment) (bool, pointArray) {
	if seg1.orientation == seg2.orientation {
		return false, pointArray{}
	}

	if seg1.orientation == horizontal {
		intersectionPoint := point {
			x: seg2.start.x,
			y: seg1.start.y,
		}

		if doesIntersectX := intersect(seg1.start.x, seg1.end.x, seg1.direction, intersectionPoint.x); !doesIntersectX {
			return false, pointArray{}
		}

		if doesIntersectY := intersect(seg2.start.y, seg2.end.y, seg2.direction, intersectionPoint.y); !doesIntersectY {
			return false, pointArray{}
		}

		return true, pointArray{intersectionPoint}
	}

	if seg1.orientation == vertical {
		intersectionPoint := point {
			x: seg1.start.x,
			y: seg2.start.y,
		}

		if doesIntersectX := intersect(seg1.start.y, seg1.end.y, seg1.direction, intersectionPoint.y); !doesIntersectX {
			return false, pointArray{}
		}

		if doesIntersectY := intersect(seg2.start.x, seg2.end.x, seg2.direction, intersectionPoint.x); !doesIntersectY {
			return false, pointArray{}
		}

		return true, pointArray{intersectionPoint}
	}

	return false, pointArray{}
}

func intersect(startCoord, endCoord, direction, intersectCoord int) bool {
	switch direction {
	case positive:
		if startCoord > intersectCoord {
			return false
		}

		if intersectCoord > endCoord {
			return false
		}

		return true

	case negative:
		if endCoord > intersectCoord {
			return false
		}

		if intersectCoord > startCoord {
			return false
		}

		return true

	default:
		return false
	}
}

func findIntersections(firstWire []lineSegment, secondWire []lineSegment) pointArray {
	var allIntersections pointArray

	for _, segment := range firstWire {
		for _, testSegment := range secondWire {
			intersect, intersectionPoints := doesIntersect(segment, testSegment)
			if intersect {
				for _, intersectionPoint := range intersectionPoints {
					// Remove the origin from list of intersections
					if intersectionPoint.x != 0 && intersectionPoint.y != 0 {
						allIntersections = append(allIntersections, intersectionPoint)
					}
				}
				break
			}
		}
	}

	return allIntersections
}

func findSmallestManhattanDistance(points pointArray) int {
	minDistance := points[0].manhattanDistance()
	for _, point := range points {
		minDistance = min(minDistance, point.manhattanDistance())
	}

	return minDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findClosestIntersection(firstWirePath string, secondWirePath string) (int, error) {
	firstWire, err := processWirePath(firstWirePath)
	if err != nil {
		return 0, err
	}

	secondWire, err := processWirePath(secondWirePath)
	if err != nil {
		return 0, err
	}

	allIntersections := findIntersections(firstWire, secondWire)
	return findSmallestManhattanDistance(allIntersections), nil
}

// Solve solves both parts of the problem
func Solve() {
	day := 3
	inputFile, err := utils.GetInput(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad things happened")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "more bad things happened")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer file.Close()

	paths, err := readWirePaths(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading wire paths failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	result, err := findClosestIntersection(paths[0], paths[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "finding closest intersection failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("found closest distance of %d\n", result)
}

func newLineSegment(start point, end point) (lineSegment, error) {
	direction, orientation := -1, -1

	if start.x == end.x {
		orientation = vertical

		if start.y <= end.y {
			direction = positive
		} else {
			direction = negative
		}

	} else {
		if start.y == end.y {
			orientation = horizontal

			if start.x <= end.x {
				direction = positive
			} else {
				direction = negative
			}
		} else {
			return lineSegment{}, errors.New("aw sausages")
		}
	}

	return lineSegment {
		start: start,
		end: end,
		orientation: orientation,
		direction: direction,
	}, nil
}
