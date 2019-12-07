package day05

import (
	"errors"
	"fmt"
	"testing"
)

func compareArrays(expected, actual []int, arrayName string) error {
	if len(actual) != len(expected) {
		return errors.New("mode lengths aren't good")
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			s := fmt.Sprintf("expected %s[%d] = %d, got %d", arrayName, i, expected[i], actual[i])
			return errors.New(s)
		}
	}

	return nil
}

// TestParseModesSymmetric tests the parseModes function in success
func TestParseModesSymmetric(t *testing.T) {
	expectedModes := []int {0, 1, 0}
	modes, err := parseModes(1, 10, 4)
	if err != nil {
		t.Error(err)
	}

	err = compareArrays(expectedModes, modes, "modes")
	if err != nil {
		t.Error(err)
	}
}

// TestParseModesAntisymmetric tests the parseModes function in success
func TestParseModesAntisymmetric(t *testing.T) {
	expectedModes := []int {1, 0, 0}
	modes, err := parseModes(1, 1, 4)
	if err != nil {
		t.Error(err)
	}

	err = compareArrays(expectedModes, modes, "modes")
	if err != nil {
		t.Error(err)
	}
}

// TestParseModesBad tests the parseModes function in failure
func TestParseModesBad(t *testing.T) {
	expectedModes := []int {1, 1, 0}
	modes, err := parseModes(1, 10, 4)
	if err != nil {
		t.Error(err)
	}

	err = compareArrays(expectedModes, modes, "modes")
	if err == nil {
		t.Error("expected error")
	}
}

// TestSplitCode tests the splitCode function
func TestSplitCode(t *testing.T) {
	modeCode, opCode := splitCode(1004)

	if modeCode != 10 {
		t.Errorf("expected modeCode = 10, got %d", modeCode)
	}

	if opCode != 4 {
		t.Errorf("expeced opCode = 4, got %d", opCode)
	}
}

// TestCreateParametersGood does the thing
func TestCreateParametersGood(t *testing.T) {
	values := []int {2, 3, 7}
	modes := []int {positionMode, positionMode, immediateMode}

	parameters, err := createParameters(values, modes)
	if err != nil {
		t.Error(err)
	}

	for i, parameter := range parameters {
		if parameter.value != values[i] {
			t.Errorf("expected parameter.value = %d, got %d", values[i], parameter.value)
		}

		if parameter.mode != modes[i] {
			t.Errorf("expected parameter.mode = %d, got %d", modes[i], parameter.mode)
		}
	}
}

// TestCreateParametersBad does the thing
func TestCreateParametersBad(t *testing.T) {
	values := []int {2, 3, 7, 2}
	modes := []int {positionMode, positionMode, immediateMode}

	_, err := createParameters(values, modes)
	if err == nil {
		t.Error("expected error")
	}
}
