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

// Instruction is used by an IntcodeComputer to perform an instruction
type Instruction struct {
	opCode int
	op func([]int, int, ...int) int
	numParams int
}

// IntcodeComputer runs an intcode program
type IntcodeComputer struct {
	initMemory []int
	memory []int
	instructionPtr int
	instructions map[int]*Instruction
	output int
}

// AddInstruction adds an Instruction to an IntcodeComputer
func (computer *IntcodeComputer) AddInstruction(in *Instruction) {
	computer.instructions[in.opCode] = in
}

func (computer *IntcodeComputer) parseParameters(numParams int) []int {
	defer computer.incrementBy(numParams)
	return computer.memory[computer.instructionPtr + 1:computer.instructionPtr + numParams]
}

func (computer *IntcodeComputer) incrementBy(value int) {
	computer.instructionPtr += value
}

func (computer *IntcodeComputer) copyMemory() {
	for i, code := range computer.initMemory {
		computer.memory[i] = code
	}
}

// Run runs the IntcodeComputer's program
func (computer *IntcodeComputer) Run() error {
	computer.copyMemory()
	computer.instructionPtr = 0
	code := computer.memory[0]

	for {
		if computer.instructionPtr >= len(computer.memory) {
			return errors.New("reached end of memory without end opcode")
		}

		if code == end {
			computer.output = computer.memory[0]
			return nil
		}

		if instruction, ok := computer.instructions[code]; ok {
			parameters := computer.parseParameters(instruction.numParams)
			result := instruction.op(computer.memory, instruction.numParams, parameters...)
			computer.memory[parameters[2]] = result
		}

		code = computer.memory[computer.instructionPtr]
	}
}

// ReadIntcode reads an intcode from a reader struct I guess
func ReadIntcode(r io.Reader) ([]int, error) {
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

// NewComputer initialises a new IntcodeComputer
func NewComputer(initMemory []int, instructions ...*Instruction) (IntcodeComputer, error) {
	computer := IntcodeComputer {
		initMemory: initMemory,
		memory: make([]int, len(initMemory)),
		instructionPtr: 0,
		output: -1,
	}

	computer.instructions = make(map[int]*Instruction, len(instructions))

	for _, instruction := range instructions {
		computer.AddInstruction(instruction)
	}

	return computer, nil
}

func instructionAdd(memory []int, numParams int, pointers ...int) int {
	sum := 0
	maxInd := numParams - 2
	for i := 0; i < maxInd; i++ {
		pointer := pointers[i]
		sum += memory[pointer]
	}
	return sum
}

func instructionMultiply(memory []int, numParams int, pointers ...int) int {
	total := 1
	maxInd := numParams - 2
	for i := 0; i < maxInd; i++ {
		pointer := pointers[i]
		total *= memory[pointer]
	}
	return total
}

var add = Instruction {
	opCode: 1,
	op: instructionAdd,
	numParams: 4,
}

var multiply = Instruction {
	opCode: 2,
	op: instructionMultiply,
	numParams: 4,
}

// PrintIntcode prints the intcode to stdout, yay!
func PrintIntcode(memory []int) {
	for _, code := range memory {
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

	initMemory, err := ReadIntcode(file)
	file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read intcode")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	initMemory[1] = 12
	initMemory[2] = 2

	computer, err := NewComputer(initMemory, &add, &multiply)
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

	fmt.Printf("replicating gravity assist program gives output = %d\n", computer.output)

	requiredOutput := 19690720

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			initMemory[1] = noun
			initMemory[2] = verb

			computer.initMemory = initMemory
			err = computer.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "computer failed to run for noun, verb = %d, %d\n", noun, verb)
				fmt.Fprintln(os.Stderr, err)
			}

			if computer.output == requiredOutput {
				fmt.Printf("output = %d when noun, verb = %d, %d\n", requiredOutput, noun, verb)
			}
		}
	}
}
