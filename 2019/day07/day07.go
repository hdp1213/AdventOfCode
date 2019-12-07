package day07

import (
	"errors"
	"fmt"
	"os"
	"github.com/hdp1213/AdventOfCode/2019/day02"
	"github.com/hdp1213/AdventOfCode/2019/day03"
	"github.com/hdp1213/AdventOfCode/2019/day05"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

// Amplifier is what we use to amplify stuff
type Amplifier struct {
	computer *day05.IntcodeComputer
	output int
}

// BaseInstructions contains all the base instructions we need
var BaseInstructions = []*day05.Instruction {
	&day05.Add,
	&day05.Multiply,
	&day05.Input,
	&day05.Output,
	&day05.JumpIfTrue,
	&day05.JumpIfFalse,
	&day05.LessThan,
	&day05.Equals,
	&day05.End,
}

// NewAmplifier makes a new amplifier struct
func NewAmplifier(code []int) (Amplifier, error) {
	computer, err := day05.NewComputer(code, BaseInstructions...)
	if err != nil {
		return Amplifier{}, err
	}

	amp := Amplifier { computer: &computer }

	return amp, nil
}

// Run does the amplifier run thing
func (amp *Amplifier) Run(phaseSetting, inputSignal int) (int, error) {
	err := amp.computer.Run(&[]int {phaseSetting, inputSignal})
	if err != nil {
		return 0, err
	}

	if len(amp.computer.Output) != 1 {
		return 0, errors.New("bad output size")
	}

	return amp.computer.Output[0], nil
}

// RunAmplification runs the full amplification gamut
func RunAmplification(code, phaseSettings []int) (int, error) {
	input := 0
	amplifier, err := NewAmplifier(code)
	if err != nil {
		return 0, err
	}

	for _, phase := range phaseSettings {
		input, err = amplifier.Run(phase, input)
		if err != nil {
			return input, err
		}
	}

	return input, nil
}

// from https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permuteArray(array []int) [][]int {
	var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }

    helper(array, len(array))
    return res
}

// Solve solves both parts of the problem
func Solve() {
	day := 7

	inputFile, err := utils.GetInput(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to get input")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to open input file")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer file.Close()

	amplifierCode, err := day02.ReadIntcode(file)
	file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read intcode")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	phaseSettingSequences := permuteArray([]int { 0, 1, 2, 3, 4 })
	outputs := map[int]int{}
	maxOutput := 0

	for i, phaseSettings := range phaseSettingSequences {
		newOutput, err := RunAmplification(amplifierCode, phaseSettings)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to run amplification for phase settings = %v\n", phaseSettings)
			fmt.Fprintln(os.Stderr, err)
			return
		}

		maxOutput = day03.Max(maxOutput, newOutput)
		outputs[i] = newOutput
	}

	for i, output := range outputs {
		if output == maxOutput {
			fmt.Printf("max output of %d achieved with phase settings = %v\n", maxOutput, phaseSettingSequences[i])
			break
		}
	}
}
