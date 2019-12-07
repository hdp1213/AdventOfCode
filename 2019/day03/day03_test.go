package day03

import (
	"fmt"
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

func TestIntersection(t *testing.T) {
	segment1, err := newLineSegment(point{x: 0, y: 0}, point{x: 10, y: 0})
	if err != nil {
		t.Error("error")
	}

	segment2, err := newLineSegment(point{x: 5, y: 10}, point{x: 5, y: 5})
	if err != nil {
		t.Error("error")
	}

	intersection := doesIntersect(segment1, segment2)
	fmt.Printf("intersection of %v and %v: %v\n", segment1, segment2, intersection)

	if intersection {
		t.Error("shit on my face you're an idiot")
	}
}
