package main

import (
	"errors"
	"fmt"
	"os"
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

func main() {
	inputFile, err := getInput(day)
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

	program := []int{1,9,10,3,2,3,11,0,99,30,40,50}
	printIntcode(program)
	runIntcode(program)
	printIntcode(program)
}