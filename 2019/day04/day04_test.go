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
