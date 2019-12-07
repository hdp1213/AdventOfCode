package day04

import (
	"testing"
)

// TestValidCodes tests for valid codes
func TestValidCodes(t *testing.T) {
	codes := []int {
		111111, // example 1
		122345,
		236678,
	}

	for _, code := range codes {
		if isValid := validate(code); !isValid {
			t.Errorf("%d should be valid", code)
		}
	}
}

// TestInvalidCodes tests for invalid codes
func TestInvalidCodes(t *testing.T) {
	codes := []int {
		223450, // example 2
		123789, // example 3
		324556,
	}

	for _, code := range codes {
		if isValid := validate(code); isValid {
			t.Errorf("%d should not be valid", code)
		}
	}
}

// TestValidCodesFurther tests for valid codes as requied in part 2
func TestValidCodesFurther(t *testing.T) {
	codes := []int {
		122345,
		236678,
		112233, // pt2, example 1
		111122, // pt2, example 2
		455669, // reddit
		455777, // reddit
		456699, // reddit
		669999, // reddit
	}

	for _, code := range codes {
		if isValid := validateFurther(code); !isValid {
			t.Errorf("%d should be valid", code)
		}
	}
}

// TestInvalidCodesFurther tests for invalid codes as requied in part 2
func TestInvalidCodesFurther(t *testing.T) {
	codes := []int {
		111111, // example 1
		223450, // example 2
		123789, // example 3
		324556,
		123444, // pt2, example 2
		455558, // reddit
		457775, // reddit
	}

	for _, code := range codes {
		if isValid := validateFurther(code); isValid {
			t.Errorf("%d should not be valid", code)
		}
	}
}
