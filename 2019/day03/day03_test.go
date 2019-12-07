package day03

import (
	"testing"
)

// func TestSmallPath(t *testing.T) {
// 	firstPath := "R8,U5,L5,D3"
// 	secondPath := "U7,R6,D4,L4"

// 	firstSegments, err := processWirePath(firstPath)
// 	if err != nil {
// 		t.Error("error")
// 	}

// 	secondSegments, err := processWirePath(secondPath)
// 	if err != nil {
// 		t.Error("error")
// 	}


// }

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
