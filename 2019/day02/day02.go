package day02

import (
	"errors"
	"fmt"
	"os"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

const (
	add = 1
	subtract = 2
	end = 99
)

const day = 2

func runIntcode(program []int) error {
	i := 0
	code := program[i]

	for {
		i++
		if i >= len(program) {
			return errors.New("reached end of program without end opcode")
		}
		
		switch code {
		case add:
			addIntcode(&program, &i)
		case subtract:
			multiplyIntcode(&program, &i)
		case end:
			return nil
		default:
			return errors.New("invalid opcode")
		}

		code = program[i]
	}
}

func addIntcode(program *[]int, i *int) {
	indices := (*program)[*i:*i+3]
	valueA, valueB := (*program)[indices[0]], (*program)[indices[1]]
	(*program)[indices[2]] = valueA + valueB
	*i += 3
}

func multiplyIntcode(program *[]int, i *int) {
	indices := (*program)[*i:*i+3]
	valueA, valueB := (*program)[indices[0]], (*program)[indices[1]]
	(*program)[indices[2]] = valueA * valueB
	*i += 3
}

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
	runIntcode(program)
	printIntcode(program)
}