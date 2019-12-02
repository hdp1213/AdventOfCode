package main


import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)


const day = 1


func fuelRequiredForMass(bah int) int {
	return (bah / 3) - 2
}


func readIntegers(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}

		result = append(result, x)
	}

	return result, scanner.Err()
}


func main() {
	inputFile, err := getInput(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad things happened")
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "more bad things happened")
	}

	defer file.Close()

	masses, err := readIntegers(file)

	total := 0
	for _, mass := range masses {
		fuel := fuelRequiredForMass(mass)
		total += fuel
	}

	fmt.Println(total)
}
