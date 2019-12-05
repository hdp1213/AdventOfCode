package day02

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

const end = 99

type instruction struct {
	opCode int
	op func([]int, int, ...int) int
	numParams int
}

type intcodeComputer struct {
	initState []int
	program []int
	instructionPtr int
	instructions map[int]*instruction
}

func (computer *intcodeComputer) addInstruction(in *instruction) {
	computer.instructions[in.opCode] = in
}

func (computer *intcodeComputer) ParseParameters(numParams int) []int {
	defer computer.IncrementBy(numParams)
	return computer.program[computer.instructionPtr + 1:computer.instructionPtr + numParams]
}

func (computer *intcodeComputer) IncrementBy(value int) {
	computer.instructionPtr += value
}

func (computer *intcodeComputer) CopyMemory() {
	for i, code := range computer.initState {
		computer.program[i] = code
	}
}

func (computer *intcodeComputer) Run() error {
	computer.CopyMemory()
	computer.instructionPtr = 0
	code := computer.program[0]

	fmt.Println("initial state:")
	printIntcode(computer.program)

	for {
		if computer.instructionPtr >= len(computer.program) {
			return errors.New("reached end of program without end opcode")
		}

		if code == end {
			fmt.Println("final state:")
			printIntcode(computer.program)
			return nil
		}

		if instruction, ok := computer.instructions[code]; ok {
			parameters := computer.ParseParameters(instruction.numParams)
			result := instruction.op(computer.program, instruction.numParams, parameters...)
			computer.program[parameters[2]] = result
		}

		code = computer.program[computer.instructionPtr]
	}
}

func readIntcode(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)
	var result []int

	for scanner.Scan() {
		t := strings.Trim(scanner.Text(), "\n ")
		x, err := strconv.Atoi(t)
		if err != nil {
			return result, err
		}

		result = append(result, x)
	}

	return result, scanner.Err()
}

func newComputer(initMemory []int, instructions ...*instruction) (intcodeComputer, error) {
	computer := intcodeComputer {
		initState: initMemory,
		program: make([]int, len(initMemory)),
		instructionPtr: 0,
	}

	computer.instructions = make(map[int]*instruction, len(instructions))

	for _, instruction := range instructions {
		computer.addInstruction(instruction)
	}

	return computer, nil
}

func instructionAdd(program []int, numParams int, pointers ...int) int {
	sum := 0
	maxInd := numParams - 2
	for i := 0; i < maxInd; i++ {
		pointer := pointers[i]
		sum += program[pointer]
	}
	return sum
}

func instructionMultiply(program []int, numParams int, pointers ...int) int {
	total := 1
	maxInd := numParams - 2
	for i := 0; i < maxInd; i++ {
		pointer := pointers[i]
		total *= program[pointer]
	}
	return total
}

func printIntcode(program []int) {
	for _, code := range program {
		fmt.Printf("%d,", code)
	}
	fmt.Println()
}

// Solve solves both parts of the problem
func Solve() {
	day := 2
	inputFile, err := utils.GetInput(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad things happened")
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "more bad things happened")
		return
	}

	defer file.Close()

	add := instruction {
		opCode: 1,
		op: instructionAdd,
		numParams: 4,
	}

	multiply := instruction {
		opCode: 2,
		op: instructionMultiply,
		numParams: 4,
	}

	initMemory, err := readIntcode(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read intcode")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	initMemory[1] = 12
	initMemory[2] = 2

	computer, err := newComputer(initMemory, &add, &multiply)
	if err != nil {
		fmt.Fprintln(os.Stderr, "computer failed to initialise")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = computer.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "computer failed to run")
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
