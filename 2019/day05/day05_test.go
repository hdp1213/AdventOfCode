package day05

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"github.com/hdp1213/AdventOfCode/2019/day02"
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

func readIntcodeString(intcode string) ([]int, error) {
	reader := strings.NewReader(intcode)
	return day02.ReadIntcode(reader)
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


type intcodeTest struct {
	intcode string
	inputs, outputs []int
}


// TestComputerEqual tests equals functionality
func TestComputerEqual(t *testing.T) {
	inputs := []int {7, 8, 9}
	outputs := []int {0, 1, 0}

	intcodeTests := []intcodeTest {
		intcodeTest {
			intcode: "3,9,8,9,10,9,4,9,99,-1,8",
			inputs: inputs,
			outputs: outputs,
		},
		intcodeTest {
			intcode: "3,3,1108,-1,8,3,4,3,99",
			inputs: inputs,
			outputs: outputs,
		},
	}

	for _, test := range intcodeTests {
		initMemory, err := readIntcodeString(test.intcode)
		if err != nil {
			t.Error(err)
		}

		computer, err := NewComputer(initMemory, &Input, &Output, &Equals, &End)
		if err != nil {
			t.Error(err)
		}

		for i, input := range test.inputs {
			err = computer.Run(input)
			if err != nil {
				t.Error(err)
			}

			if len(computer.output) != 1 {
				t.Errorf("expected len(computer.output) == 1, got %d", len(computer.output))
			}

			if computer.output[0] != test.outputs[i] {
				t.Errorf("expected output = %d, got %d", test.outputs[i], computer.output[0])
			}
		}
	}
}

// TestComputerLessThan tests less than functionality
func TestComputerLessThan(t *testing.T) {
	inputs := []int {7, 8, 9}
	outputs := []int {1, 0, 0}

	intcodeTests := []intcodeTest {
		intcodeTest {
			intcode: "3,9,7,9,10,9,4,9,99,-1,8",
			inputs: inputs,
			outputs: outputs,
		},
		intcodeTest {
			intcode: "3,3,1107,-1,8,3,4,3,99",
			inputs: inputs,
			outputs: outputs,
		},
	}

	for _, test := range intcodeTests {
		initMemory, err := readIntcodeString(test.intcode)
		if err != nil {
			t.Error(err)
		}

		computer, err := NewComputer(initMemory, &Input, &Output, &LessThan, &End)
		if err != nil {
			t.Error(err)
		}

		for i, input := range test.inputs {
			err =computer.Run(input)
			if err != nil {
				t.Error(err)
			}

			if len(computer.output) != 1 {
				t.Errorf("expected len(computer.output) == 1, got %d", len(computer.output))
			}

			if computer.output[0] != test.outputs[i] {
				t.Errorf("expected output = %d, got %d", test.outputs[i], computer.output[0])
			}
		}
	}
}

// TestComputerJump tests jump functionality
func TestComputerJump(t *testing.T) {
	inputs := []int {0, 1, -1}
	outputs := []int {0, 1, 1}

	intcodeTests := []intcodeTest {
		intcodeTest {
			intcode: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			inputs: inputs,
			outputs: outputs,
		},
		intcodeTest {
			intcode: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			inputs: inputs,
			outputs: outputs,
		},
	}

	for _, test := range intcodeTests {
		initMemory, err := readIntcodeString(test.intcode)
		if err != nil {
			t.Error(err)
		}

		computer, err := NewComputer(initMemory, &Add, &Input, &Output, &JumpIfTrue, &JumpIfFalse, &End)
		if err != nil {
			t.Error(err)
		}

		for i, input := range test.inputs {
			err = computer.Run(input)
			if err != nil {
				t.Error(err)
			}

			if len(computer.output) != 1 {
				t.Errorf("expected len(computer.output) == 1, got %d", len(computer.output))
			}

			if computer.output[0] != test.outputs[i] {
				t.Errorf("expected output = %d, got %d", test.outputs[i], computer.output[0])
			}
		}
	}
}

// TestComputerLargerExample tests the larger part 2 example
func TestComputerLargerExample(t *testing.T) {
	test := intcodeTest {
		intcode: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		inputs: []int {-100, 7, 8, 9, 100},
		outputs: []int {999, 999, 1000, 1001, 1001},
	}

	initMemory, err := readIntcodeString(test.intcode)
	if err != nil {
		t.Error(err)
	}

	computer, err := NewComputer(initMemory, &Add, &Multiply, &Input, &Output, &JumpIfTrue, &JumpIfFalse, &LessThan, &Equals, &End)
	if err != nil {
		t.Error(err)
	}

	for i, input := range test.inputs {
		err = computer.Run(input)
		if err != nil {
			t.Error(err)
		}

		if len(computer.output) != 1 {
			t.Errorf("expected len(computer.output) == 1, got %d", len(computer.output))
		}

		if computer.output[0] != test.outputs[i] {
			t.Errorf("expected output = %d, got %d", test.outputs[i], computer.output[0])
		}
	}
}
