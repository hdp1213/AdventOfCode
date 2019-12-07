package day04

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

type code int

func validate(code int) bool {
	factor := 10
	div := code

	pastDigit := -1
	hasSameDigit := false

	for div > 0 {
		div /= factor
		digit := code - div * factor

		if digit > pastDigit && pastDigit != -1 {
			return false
		}

		if digit == pastDigit && !hasSameDigit {
			hasSameDigit = true
		}

		pastDigit = digit
		code = div
	}

	return hasSameDigit
}

func testValidation(code int) func(int) bool {
	// state := code
	return func(x int) bool {
		return true
	}
}

func readRange(r io.Reader) (int, int, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return -1, -1, err
	}

	dataString := strings.TrimSpace(string(data[:len(data)]))
	codeRange := strings.SplitN(dataString, "-", 2)

	a, err := strconv.Atoi(codeRange[0])
	if err != nil {
		return -1, -1, err
	}

	b, err := strconv.Atoi(codeRange[1])
	if err != nil {
		return -1, -1, err
	}

	return a, b, nil
}

// Solve solves both parts of the problem
func Solve() {
	day := 4

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

	a, b, err := readRange(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read range")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("%d %d\n", a, b)
}
