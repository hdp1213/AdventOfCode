package day02

import (
	"errors"
	"fmt"
	"os"
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

func (computer *intcodeComputer) Run() error {
	computer.instructionPtr = 0
	code := computer.program[0]

	for {
		if computer.instructionPtr >= len(computer.program) {
			return errors.New("reached end of program without end opcode")
		}

		if code == end {
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

func newComputer(program []int, instructions ...*instruction) intcodeComputer {
	computer := intcodeComputer {
		program: program,
		instructionPtr: 0,
	}

	computer.instructions = make(map[int]*instruction, len(instructions))

	for _, instruction := range instructions {
		computer.addInstruction(instruction)
	}

	return computer
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

const day = 2

func printIntcode(program []int) {
	for _, code := range program {
		fmt.Printf("%d,", code)
	}
	fmt.Println()
}

// Solve solves both parts of the problem
func Solve() {
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

	program := []int{1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,5,19,23,1,23,6,27,1,5,27,31,1,31,6,35,1,9,35,39,2,10,39,43,1,43,6,47,2,6,47,51,1,5,51,55,1,55,13,59,1,59,10,63,2,10,63,67,1,9,67,71,2,6,71,75,1,5,75,79,2,79,13,83,1,83,5,87,1,87,9,91,1,5,91,95,1,5,95,99,1,99,13,103,1,10,103,107,1,107,9,111,1,6,111,115,2,115,13,119,1,10,119,123,2,123,6,127,1,5,127,131,1,5,131,135,1,135,6,139,2,139,10,143,2,143,9,147,1,147,6,151,1,151,13,155,2,155,9,159,1,6,159,163,1,5,163,167,1,5,167,171,1,10,171,175,1,13,175,179,1,179,2,183,1,9,183,0,99,2,14,0,0}
	printIntcode(program)

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

	computer := newComputer(program, &add, &multiply)
	computer.Run()
	printIntcode(computer.program)
}
