package day03

import (
	"testing"
)

// TestSmallPaths does an in-depth test of the worked example
func TestSmallPaths(t *testing.T) {
	firstPath := "R8,U5,L5,D3"
	secondPath := "U7,R6,D4,L4"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 2 {
		t.Errorf("expected len(intersections) == 2, got %d", len(intersections))
	}

	point1 := point {x: 3, y: 3}
	point2 := point {x: 6, y: 5}

	if !intersections.contains(point1) {
		t.Errorf("could not find point %v", point1)
	}

	if !intersections.contains(point2) {
		t.Errorf("could not find point %v", point2)
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 6 {
		t.Errorf("expected distance == 6, got %d", distance)
	}

	wireDistance := findSmallestWireDistance(firstWire, secondWire, intersections)

	if wireDistance != 30 {
		t.Errorf("expected wireDistance == 30, got %d", wireDistance)
	}
}

// TestSmallPathsFlipped90 is the same as TestSmallPaths but flipped 90 degrees
func TestSmallPathsFlipped90(t *testing.T) {
	firstPath := "U8,L5,D5,R3"
	secondPath := "L7,U6,R4,D4"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 2 {
		t.Errorf("expected len(intersections) == 2, got %d", len(intersections))
	}

	point1 := point {x: -3, y: 3}
	point2 := point {x: -5, y: 6}

	if !intersections.contains(point1) {
		t.Errorf("could not find point %v", point1)
	}

	if !intersections.contains(point2) {
		t.Errorf("could not find point %v", point2)
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 6 {
		t.Errorf("expected distance == 6, got %d", distance)
	}
}

// TestSmallPathsFlipped180 is the same as TestSmallPaths but flipped 180 degrees
func TestSmallPathsFlipped180(t *testing.T) {
	firstPath := "L8,D5,R5,U3"
	secondPath := "D7,L6,U4,R4"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 2 {
		t.Errorf("expected len(intersections) == 2, got %d", len(intersections))
	}

	point1 := point {x: -3, y: -3}
	point2 := point {x: -6, y: -5}

	if !intersections.contains(point1) {
		t.Errorf("could not find point %v", point1)
	}

	if !intersections.contains(point2) {
		t.Errorf("could not find point %v", point2)
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 6 {
		t.Errorf("expected distance == 6, got %d", distance)
	}
}

// TestSmallPathsFlipped270 is the same as TestSmallPaths but flipped 270 degrees
func TestSmallPathsFlipped270(t *testing.T) {
	firstPath := "D8,R5,U5,L3"
	secondPath := "R7,D6,L4,U4"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 2 {
		t.Errorf("expected len(intersections) == 2, got %d", len(intersections))
	}

	point1 := point {x: 3, y: -3}
	point2 := point {x: 5, y: -6}

	if !intersections.contains(point1) {
		t.Errorf("could not find point %v", point1)
	}

	if !intersections.contains(point2) {
		t.Errorf("could not find point %v", point2)
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 6 {
		t.Errorf("expected distance == 6, got %d", distance)
	}
}

// TestFirstMediumPaths tests the first non-worked example
func TestFirstMediumPaths(t *testing.T) {
	firstPath := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	secondPath := "U62,R66,U55,R34,D71,R55,D58,R83"

	distance, err := findClosestManhattanIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if distance != 159 {
		t.Errorf("expected distance == 159, got %d", distance)
	}

	wireDistance, err := findClosestWireIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if wireDistance != 610 {
		t.Errorf("expected wireDistance == 610, got %d", wireDistance)
	}
}

// TestSecondMediumPaths tests the second non-worked example
func TestSecondMediumPaths(t *testing.T) {
	firstPath := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	secondPath := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	distance, err := findClosestManhattanIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if distance != 135 {
		t.Errorf("expected distance == 135, got %d", distance)
	}

	wireDistance, err := findClosestWireIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if wireDistance != 410 {
		t.Errorf("expected wireDistance == 410, got %d", wireDistance)
	}
}

// TestMultipleIntersections tests location of multiple intersections along one line segment
func TestMultipleIntersections(t *testing.T) {
	firstPath := "U5,R50"
	secondPath := "R10,U10,R10,D10,R10,U10,R20"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	allIntersections := findIntersections(firstWire, secondWire)

	if len(allIntersections) != 3 {
		t.Errorf("expected len(allIntersections) == 3, got %d", len(allIntersections))
	}
}

// TestElbowIntersection tests intersection of two wires touching by only their bend
func TestElbowIntersection(t *testing.T) {
	firstPath := "R5,U10,R10"
	secondPath := "U10,R5,U2"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 1 {
		t.Errorf("expected len(intersections) == 1, got %d", len(intersections))
	}

	point1 := point {x: 5, y: 10}

	if !intersections.contains(point1) {
		t.Errorf("could not find point %v", point1)
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 15 {
		t.Errorf("expected distance == 15, got %d", distance)
	}
}

// TestOverlapIntersection tests non-standard case of two wires overlapping
func TestOverlapIntersection(t *testing.T) {
	firstPath := "D2,R10,U7"
	secondPath := "R10,U2,R2"

	firstWire, err := processWirePath(firstPath)
	if err != nil {
		t.Error(err)
	}

	secondWire, err := processWirePath(secondPath)
	if err != nil {
		t.Error(err)
	}

	intersections := findIntersections(firstWire, secondWire)

	if len(intersections) != 3 {
		t.Errorf("expected len(intersections) == 3, got %d", len(intersections))
	}

	testPoints := pointArray{
		point {x: 10, y: 0},
		point {x: 10, y: 1},
		point {x: 10, y: 2},
	}

	for _, point := range testPoints {
		if !intersections.contains(point) {
			t.Errorf("could not find point %v", point)
		}
	}

	distance := findSmallestManhattanDistance(intersections)

	if distance != 10 {
		t.Errorf("expected distance == 10, got %d", distance)
	}
}

// TestNoIntersectionHorizontal tests non-intersection of two line segments, first horizontal
func TestNoIntersectionHorizontal(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 10, y: 0})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: 5}, point{x: 5, y: 10})
	if err != nil {
		t.Error("error")
	}

	intersection, points := doesIntersect(segment1, segment2)

	if intersection {
		t.Error("segments should not intersect")
	}

	if len(points) != 0 {
		t.Error("intersection point is non-zero")
	}
}

// TestIntersectionHorizontal tests intersection of two line segments, first horizontal
func TestIntersectionHorizontal(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 10, y: 0})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: -10}, point{x: 5, y: 10})
	if err != nil {
		t.Error("error")
	}

	intersection, points := doesIntersect(segment1, segment2)

	if !intersection {
		t.Error("segments should intersect")
	}

	point := points[0]

	if point.x != 5 && point.y != 0 {
		t.Error("intersection point should be {5, 0}")
	}
}

// TestNoIntersectionVertical tests non-intersection of two line segments, first vertical
func TestNoIntersectionVertical(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 0, y: 10})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: 5}, point{x: 10, y: 5})
	if err != nil {
		t.Error("error")
	}

	intersection, points := doesIntersect(segment1, segment2)

	if intersection {
		t.Error("segments should not intersect")
	}

	if len(points) != 0 {
		t.Error("intersection point is non-zero")
	}
}

// TestIntersectionVertical tests intersection of two line segments, first vertical
func TestIntersectionVertical(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 0, y: 10})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: -10, y: 5}, point{x: 10, y: 5})
	if err != nil {
		t.Error("error")
	}

	intersection, points := doesIntersect(segment1, segment2)

	if !intersection {
		t.Error("segments should intersect")
	}

	point := points[0]

	if point.x != 0 && point.y != 5 {
		t.Error("intersection point should be {0, 5}")
	}
}

// TestGetDistanceExact tests getting wire distance to the point where the wire ends
func TestGetDistanceExact(t *testing.T) {
	path := "R4,U3"
	point := point{x: 4, y: 3}

	wire, err := processWirePath(path)
	if err != nil {
		t.Error(err)
	}

	wireDistance := getDistanceToPoint(wire, point)

	if wireDistance != 7 {
		t.Errorf("expected wireDistance == 7, got %d", wireDistance)
	}
}

// TestGetDistancePartial tests getting wire distance to a point on the last segment
func TestGetDistancePartial(t *testing.T) {
	path := "R4,U10"
	point := point{x: 4, y: 3}

	wire, err := processWirePath(path)
	if err != nil {
		t.Error(err)
	}

	wireDistance := getDistanceToPoint(wire, point)

	if wireDistance != 7 {
		t.Errorf("expected wireDistance == 7, got %d", wireDistance)
	}
}

// TestGetDistanceOvershoot tests getting wire distance to a point on a segment that is not the last
func TestGetDistanceOvershoot(t *testing.T) {
	path := "R4,U10,L3,D2"
	point := point{x: 4, y: 3}

	wire, err := processWirePath(path)
	if err != nil {
		t.Error(err)
	}

	wireDistance := getDistanceToPoint(wire, point)

	if wireDistance != 7 {
		t.Errorf("expected wireDistance == 7, got %d", wireDistance)
	}
}
