package day05

import (
	"errors"
	"fmt"
	"os"
	"github.com/hdp1213/AdventOfCode/2019/day02"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

// EndCode is the opcode used to terminate an intcode program
const EndCode = 99

const (
	positionMode = 0
	immediateMode = 1
)

// Parameter is what is parsed by that function
type Parameter struct {
	value, mode int
}

func (parameter *Parameter) getValue(memory []int) int {
	switch parameter.mode {
	case positionMode:
		return memory[parameter.value]
	case immediateMode:
		return parameter.value
	default:
		return 0
	}
}

func createParameters(values, modes []int) ([]Parameter, error) {
	parameters := []Parameter{}

	if len(values) != len(modes) {
		return parameters, errors.New("values and modes are not the same length")
	}

	for i := 0; i < len(values); i++ {
		parameters = append(parameters, Parameter {value: values[i], mode: modes[i]})
	}

	return parameters, nil
}

// Instruction is used by an IntcodeComputer to perform an instruction
type Instruction struct {
	opCode, numValues int
	op func(*IntcodeComputer, int, []Parameter) error
	doesJump bool
}

// Execute runs an instruction's operation on the given computer
func (instruction *Instruction) Execute(computer *IntcodeComputer, input int, parameters ...Parameter) error {
	if len(parameters) != instruction.numValues - 1 {
		errorMessage := fmt.Sprintf("instruction opcode %d incorrectly configured: bad numValues", instruction.opCode)
		return errors.New(errorMessage)
	}

	err := instruction.op(computer, input, parameters)
	if err != nil {
		errorMessage := fmt.Sprintf("opcode %d failed: %v", instruction.opCode, err)
		return errors.New(errorMessage)
	}

	return nil
}

func instructionAdd(computer *IntcodeComputer, input int, parameters []Parameter) error {
	if parameters[2].mode == immediateMode {
		return errors.New("write parameter cannot be in immediate mode")
	}

	sum := parameters[0].getValue(computer.memory) + parameters[1].getValue(computer.memory)
	computer.memory[parameters[2].value] = sum
	return nil
}

func instructionMultiply(computer *IntcodeComputer, input int, parameters []Parameter) error {
	if parameters[2].mode == immediateMode {
		return errors.New("write parameter cannot be in immediate mode")
	}

	product := parameters[0].getValue(computer.memory) * parameters[1].getValue(computer.memory)
	computer.memory[parameters[2].value] = product
	return nil
}

func instructionInput(computer *IntcodeComputer, input int, parameters []Parameter) error {
	if parameters[0].mode == immediateMode {
		return errors.New("write parameter cannot be in immediate mode")
	}

	computer.memory[parameters[0].value] = input
	return nil
}

func instructionOutput(computer *IntcodeComputer, input int, parameters []Parameter) error {
	output := parameters[0].getValue(computer.memory)
	fmt.Println(output)
	return nil
}

func instructionJumpIfTrue(computer *IntcodeComputer, input int, parameters []Parameter) error {
	if parameters[0].getValue(computer.memory) != 0 {
		// set instruction pointer to value of second parameter
	}

	return nil
}

func instructionJumpIfFalse(computer *IntcodeComputer, input int, parameters []Parameter) error {
	return nil
}

func instructionLessThan(computer *IntcodeComputer, input int, parameters []Parameter) error {
	return nil
}

func instructionEquals(computer *IntcodeComputer, input int, parameters []Parameter) error {
	return nil
}

func instructionEnd(computer *IntcodeComputer, input int, parameters []Parameter) error {
	return nil
}

// Add is an instruction that adds two elements together
var Add = Instruction {
	opCode: 1,
	op: instructionAdd,
	numValues: 4,
	doesJump: false,
}

// Multiply is an instruction that multiplies two elements together
var Multiply = Instruction {
	opCode: 2,
	op: instructionMultiply,
	numValues: 4,
	doesJump: false,
}

// Input is an instruction that takes input and saves it to a memory location
var Input = Instruction {
	opCode: 3,
	op: instructionInput,
	numValues: 2,
	doesJump: false,
}

// Output is an instruction that writes output from given memory location
var Output = Instruction {
	opCode: 4,
	op: instructionOutput,
	numValues: 2,
	doesJump: false,
}

// End is an instruction that ends the program
var End = Instruction {
	opCode: EndCode,
	op: instructionEnd,
	numValues: 1,
	doesJump: false,
}

// IntcodeComputer runs an intcode program
type IntcodeComputer struct {
	initMemory, memory []int
	instructionPtr int
	instructions map[int]*Instruction
	output []int
}

// AddInstruction adds an Instruction to an IntcodeComputer
func (computer *IntcodeComputer) AddInstruction(in *Instruction) {
	computer.instructions[in.opCode] = in
}

func (computer *IntcodeComputer) parseParameters(numValues int, doIncrement bool) []int {
	if doIncrement {
		defer computer.incrementBy(numValues)
	}
	return computer.memory[computer.instructionPtr + 1:computer.instructionPtr + numValues]
}

func (computer *IntcodeComputer) incrementBy(value int) {
	computer.instructionPtr += value
}

func (computer *IntcodeComputer) copyMemory() {
	for i, code := range computer.initMemory {
		computer.memory[i] = code
	}
}

func (computer *IntcodeComputer) parseNextInstruction() (*Instruction, []Parameter, error) {
	fullCode := computer.memory[computer.instructionPtr]
	modeCode, opCode := splitCode(fullCode)

	if instruction, ok := computer.instructions[opCode]; ok {
		values := computer.parseParameters(instruction.numValues, !instruction.doesJump)
		modes, err := parseModes(opCode, modeCode, instruction.numValues)
		if err != nil {
			return instruction, []Parameter{}, err
		}

		parameters, err := createParameters(values, modes)
		if err != nil {
			return instruction, parameters, err
		}

		return instruction, parameters, nil
	}

	errorMessage := fmt.Sprintf("opcode %d not supported", opCode)
	return &Instruction{}, []Parameter{}, errors.New(errorMessage)
}

func splitCode(fullCode int) (modeCode, opCode int) {
	const modeFactor = 100
	modeCode = fullCode / modeFactor
	opCode = fullCode - modeCode * modeFactor
	return
}

func parseModes(opCode, modeCode, numValues int) ([]int, error) {
	const factor = 10
	modes := make([]int, numValues - 1)
	value := modeCode
	i := 0

	for value > 0 {
		if i == len(modes) {
			errorMessage := fmt.Sprintf("too many modes specified for opcode %d", opCode)
			return modes, errors.New(errorMessage)
		}

		digit := value - (value / factor) * factor
		modes[i] = digit
		value /= factor
		i++
	}

	return modes, nil
}

// Run runs the IntcodeComputer's program
func (computer *IntcodeComputer) Run(input int) error {
	computer.copyMemory()
	computer.instructionPtr = 0

	instruction, parameters, err := computer.parseNextInstruction()
	if err != nil {
		return err
	}

	for {
		if instruction.opCode == EndCode {
			return nil
		}

		instruction.Execute(computer, input, parameters...)

		instruction, parameters, err = computer.parseNextInstruction()
		if err != nil {
			return err
		}

		if computer.instructionPtr >= len(computer.memory) {
			return errors.New("reached end of memory without end opcode")
		}
	}
}

// NewComputer initialises a new IntcodeComputer
func NewComputer(initMemory []int, instructions ...*Instruction) (IntcodeComputer, error) {
	computer := IntcodeComputer {
		initMemory: initMemory,
		memory: make([]int, len(initMemory)),
		instructionPtr: 0,
		output: []int{},
	}

	computer.instructions = make(map[int]*Instruction, len(instructions))

	for _, instruction := range instructions {
		computer.AddInstruction(instruction)
	}

	return computer, nil
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
	day := 5

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

	initMemory, err := day02.ReadIntcode(file)
	file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read intcode")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	computer, err := NewComputer(initMemory, &Add, &Multiply, &Input, &Output, &End)
	if err != nil {
		fmt.Fprintln(os.Stderr, "computer failed to initialise")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = computer.Run(1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "computer failed to run, reached instruction %d\n", computer.instructionPtr)
		// fmt.Fprintln(os.Stderr, "memory dump:")
		// PrintIntcode(computer.memory)
		return 
	}
}
