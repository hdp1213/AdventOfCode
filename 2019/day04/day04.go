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

func validate(code int) bool {
	return validateRecursive(code, 10, false)
}

func validateRecursive(x, prevDigit int, hasSameDigit bool) bool {
	const factor = 10
	div := x / factor
	digit := x - div * factor

	if digit > prevDigit {
		return false
	}

	if digit == prevDigit && !hasSameDigit {
		hasSameDigit = true
	}

	if div > 0 {
		return validateRecursive(div, digit, hasSameDigit)
	}

	return hasSameDigit
}

func validateFurther(code int) bool {
	return validateFurtherRecursive(code, 10, map[int]int{})
}

func validateFurtherRecursive(x, prevDigit int, groupSizes map[int]int) bool {
	const factor = 10
	div := x / factor
	digit := x - div * factor

	if digit > prevDigit {
		return false
	}

	if digit == prevDigit {
		if _, ok := groupSizes[digit]; ok {
			groupSizes[digit]++
		} else {
			groupSizes[digit] = 2
		}
	}

	if div > 0 {
		return validateFurtherRecursive(div, digit, groupSizes)
	}

	for _, count := range groupSizes {
		if count == 2 {
			return true
		}
	}

	return false
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

	totalValidCodes, totalFurtherValidCodes := 0, 0
	for code := a; code <= b; code++ {
		if validate(code) {
			totalValidCodes++
		}

		if validateFurther(code) {
			totalFurtherValidCodes++
		}
	}

	fmt.Printf("total valid codes found in [%d,%d] = %d\n", a, b, totalValidCodes)
	fmt.Printf("total further valid codes found in [%d,%d] = %d\n", a, b, totalFurtherValidCodes)
}
