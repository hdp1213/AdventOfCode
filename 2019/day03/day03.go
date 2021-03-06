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
	x, y int
}

func (point *point) manhattanDistance() int {
	return abs(point.x) + abs(point.y)
}

type pointArray []point

func (points pointArray) contains(testPoint point) bool {
	for _, point := range points {
		if testPoint.x == point.x && testPoint.y == point.y {
			return true
		}
	}
	return false
}

type lineSegment struct {
	start, end point
	orientation, direction int
}

func (segment lineSegment) length() int {
	return segment.distanceTo(segment.end)
}

func (segment lineSegment) distanceTo(point point) int {
	switch segment.orientation {
	case horizontal:
		return abs(segment.start.x - point.x)
	case vertical:
		return abs(segment.start.y - point.y)
	default:
		return -1
	}
}

func newLineSegment(start, end point) (lineSegment, error) {
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
			return lineSegment{}, errors.New("Failed to make new line segment")
		}
	}

	return lineSegment {
		start: start,
		end: end,
		orientation: orientation,
		direction: direction,
	}, nil
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

func doesIntersectPoint(segment lineSegment, point point) bool {
	if segment.orientation == horizontal {
		return intersect(segment.start.y, segment.end.y, segment.direction, point.y)
	}

	if segment.orientation == vertical {
		return intersect(segment.start.x, segment.end.x, segment.direction, point.x)
	}

	return false
}

func doesIntersect(seg1, seg2 lineSegment) (bool, pointArray) {
	if seg1.orientation == seg2.orientation {
		var intersectionPoints pointArray

		if seg1.orientation == horizontal && seg1.start.y == seg2.start.y {
			minX, maxX := min(seg1.start.x, seg1.end.x), Max(seg1.start.x, seg1.end.x)

			for x := minX; x <= maxX; x++ {
				if intersect(seg2.start.x, seg2.end.x, seg2.direction, x) {
					intersectionPoints = append(intersectionPoints, point{x: x, y: seg1.start.y})
				}
			}

			if len(intersectionPoints) > 0 {
				return true, intersectionPoints
			}

			return false, pointArray{}
		}

		if seg1.orientation == vertical && seg1.start.x == seg2.start.x {
			minY, maxY := min(seg1.start.y, seg1.end.y), Max(seg1.start.y, seg1.end.y)

			for y := minY; y <= maxY; y++ {
				if intersect(seg2.start.y, seg2.end.y, seg2.direction, y) {
					intersectionPoints = append(intersectionPoints, point{x: seg1.start.x, y: y})
				}
			}

			if len(intersectionPoints) > 0 {
				return true, intersectionPoints
			}

			return false, pointArray{}
		}

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

func findIntersections(firstWire, secondWire []lineSegment) pointArray {
	var allIntersections pointArray

	for _, segment := range firstWire {
		for _, testSegment := range secondWire {
			intersect, intersectionPoints := doesIntersect(segment, testSegment)
			if intersect {
				// Ignore origin intersection
				if len(intersectionPoints) == 1 && intersectionPoints[0].x == 0 && intersectionPoints[0].y == 0 {
					continue
				}

				for _, intersectionPoint := range intersectionPoints {
					// Don't include duplicates
					if !allIntersections.contains(intersectionPoint) {
						allIntersections = append(allIntersections, intersectionPoint)
					}
				}
			}
		}
	}

	return allIntersections
}

func findSmallestManhattanDistance(points pointArray) int {
	if len(points) == 0 {
		return -1
	}

	minDistance := points[0].manhattanDistance()
	for _, point := range points {
		minDistance = min(minDistance, point.manhattanDistance())
	}

	return minDistance
}

func findSmallestWireDistance(firstWire, secondWire []lineSegment, intersections pointArray) int {
	if len(intersections) == 0 {
		return -1
	}

	minDistance := getDistanceToPoint(firstWire, intersections[0]) + getDistanceToPoint(secondWire, intersections[0])
	for _, point := range intersections {
		newDistance := getDistanceToPoint(firstWire, point) + getDistanceToPoint(secondWire, point)
		minDistance = min(minDistance, newDistance)
	}

	return minDistance
}

func getDistanceToPoint(wire []lineSegment, point point) int {
	totalDistance := 0
	for _, segment := range wire {
		if doesIntersectPoint(segment, point) {
			totalDistance += segment.distanceTo(point)
			break
		} else {
			totalDistance += segment.length()
		}
	}

	return totalDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max gets the maximum value of two integers
func Max(a, b int) int {
	if a > b {
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

func findClosestManhattanIntersection(firstWirePath, secondWirePath string) (int, error) {
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

func findClosestWireIntersection(firstWirePath, secondWirePath string) (int, error) {
	firstWire, err := processWirePath(firstWirePath)
	if err != nil {
		return 0, err
	}

	secondWire, err := processWirePath(secondWirePath)
	if err != nil {
		return 0, err
	}

	allIntersections := findIntersections(firstWire, secondWire)
	return findSmallestWireDistance(firstWire, secondWire, allIntersections), nil
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

	result, err := findClosestManhattanIntersection(paths[0], paths[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "finding closest manhattan intersection failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("found closest distance of %d\n", result)

	wireDistance, err := findClosestWireIntersection(paths[0], paths[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "finding closest wire intersection failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("found closest wire distance of %d\n", wireDistance)
}
