package day03

import (
	"testing"
)

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
		t.Errorf("expected distance == 6, got %d", point2)
	}
}

func TestFirstMediumPaths(t *testing.T) {
	firstPath := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	secondPath := "U62,R66,U55,R34,D71,R55,D58,R83"

	distance, err := findClosestIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if distance != 159 {
		t.Errorf("expected distance == 159, got %d", distance)
	}
}

func TestSecondMediumPaths(t *testing.T) {
	firstPath := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	secondPath := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	distance, err := findClosestIntersection(firstPath, secondPath)
	if err != nil {
		t.Error(err)
	}

	if distance != 135 {
		t.Errorf("expected distance == 135, got %d", distance)
	}
}

func TestNoIntersectionHorizontal(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 10, y: 0})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: 5}, point{x: 5, y: 10})
	if err != nil {
		t.Error("error")
	}

	intersection, pt := doesIntersect(segment1, segment2)

	if intersection {
		t.Error("segments should not intersect")
	}

	if pt.x != 0 && pt.y != 0 {
		t.Error("intersection point is non-zero")
	}
}

func TestIntersectionHorizontal(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 10, y: 0})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: -10}, point{x: 5, y: 10})
	if err != nil {
		t.Error("error")
	}

	intersection, pt := doesIntersect(segment1, segment2)

	if !intersection {
		t.Error("segments should intersect")
	}

	if pt.x != 5 && pt.y != 0 {
		t.Error("intersection point should be {5, 0}")
	}
}

func TestNoIntersectionVertical(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 0, y: 10})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: 5}, point{x: 10, y: 5})
	if err != nil {
		t.Error("error")
	}

	intersection, pt := doesIntersect(segment1, segment2)

	if intersection {
		t.Error("segments should not intersect")
	}

	if pt.x != 0 && pt.y != 0 {
		t.Error("intersection point is non-zero")
	}
}

func TestIntersectionVertical(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 0, y: 10})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: -10, y: 5}, point{x: 10, y: 5})
	if err != nil {
		t.Error("error")
	}

	intersection, pt := doesIntersect(segment1, segment2)

	if !intersection {
		t.Error("segments should intersect")
	}

	if pt.x != 0 && pt.y != 5 {
		t.Error("intersection point should be {0, 5}")
	}
}
