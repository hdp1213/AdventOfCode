package day07

import (
	"strings"
	"testing"
	"github.com/hdp1213/AdventOfCode/2019/day02"
)

func runAmplificationFromString(rawCode string, phaseSettings []int) (int, error) {
	r := strings.NewReader(rawCode)
	code, err := day02.ReadIntcode(r)
	if err != nil {
		return 0, err
	}

	output, err := RunAmplification(code, phaseSettings)
	if err != nil {
		return output, err
	}

	return output, nil
}

// TestFirstExample does the first example test
func TestFirstExample(t *testing.T) {
	code := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	phaseSettings := []int { 4, 3, 2, 1, 0 }
	expectedOutput := 43210

	output, err := runAmplificationFromString(code, phaseSettings)
	if err != nil {
		t.Error(err)
	}

	if output != expectedOutput {
		t.Errorf("expected output = %d, got %d", expectedOutput, output)
	}
}

// TestSecondExample does the second example test
func TestSecondExample(t *testing.T) {
	code := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	phaseSettings := []int { 0, 1, 2, 3, 4 }
	expectedOutput := 54321

	output, err := runAmplificationFromString(code, phaseSettings)
	if err != nil {
		t.Error(err)
	}

	if output != expectedOutput {
		t.Errorf("expected output = %d, got %d", expectedOutput, output)
	}
}

// TestThirdExample does the third example test
func TestThirdExample(t *testing.T) {
	code := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	phaseSettings := []int { 1, 0, 4, 3, 2 }
	expectedOutput := 65210

	output, err := runAmplificationFromString(code, phaseSettings)
	if err != nil {
		t.Error(err)
	}

	if output != expectedOutput {
		t.Errorf("expected output = %d, got %d", expectedOutput, output)
	}
}
