package main


import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)


const day = 1


func fuelRequiredForMass(mass int) int {
	return (mass / 3) - 2
}

func fuelRequiredForMassAndFuel(mass int) int {
	fuel := fuelRequiredForMass(mass)
	if fuel < 0 {
		return 0
	}
	return fuel + fuelRequiredForMassAndFuel(fuel)
}

func fuelRequiredForEverything(moduleMasses []int) int {
	totalFuel := 0
	for _, moduleMass := range moduleMasses {
		totalFuel += fuelRequiredForMassAndFuel(moduleMass)
	}
	return totalFuel
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
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "more bad things happened")
		return
	}

	defer file.Close()

	moduleMasses, err := readIntegers(file)
	total := fuelRequiredForEverything(moduleMasses)

	fmt.Println(total)
}
