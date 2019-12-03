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

func mapAndSum(array []int, f func(int) int) int {
	total := 0
	for _, elem := range array {
		total += f(elem)
	}
	return total
}

func fuelRequiredForModules(moduleMasses []int) int {
	return mapAndSum(moduleMasses, fuelRequiredForMass)
}

func fuelRequiredForEverything(moduleMasses []int) int {
	return mapAndSum(moduleMasses, fuelRequiredForMassAndFuel)
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
	fuelForModules := fuelRequiredForModules(moduleMasses)
	totalFuel := fuelRequiredForEverything(moduleMasses)

	fmt.Printf("Fuel for only modules:   %d\n", fuelForModules)
	fmt.Printf("Fuel for modules + fuel: %d\n", totalFuel)
}
