package day06

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "sort"
	"strings"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)


// COM thing
const COM = "COM"

type orbit struct {
	name string
	parent *orbit
}

type orbitList []orbit

type preOrbitList map[string][]string

func (orbits *orbitList) get(name string) (orbit, bool) {
	for _, orbit := range *orbits {
		if orbit.name == name {
			return orbit, true
		}
	}
	return orbit{}, false
}

func readPreOrbits(r io.Reader) (preOrbitList, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	orbits := preOrbitList{}

	for scanner.Scan() {
		line := scanner.Text()
		orbitBodies := strings.SplitN(line, ")", 2)
		parent, child := orbitBodies[0], orbitBodies[1]

		if _, ok := orbits[parent]; ok {
			orbits[parent] = append(orbits[parent], child)
		} else {
			orbits[parent] = []string { child }
		}
	}

	return orbits, scanner.Err()
}

func processOrbits(orbits preOrbitList) orbitList {
	newOrbits := orbitList{
		orbit { name: COM, parent: nil },
	}

	return *addChildOrbitsOf(COM, orbits, &newOrbits)
}

func addChildOrbitsOf(parentName string, orbits preOrbitList, fullOrbits *orbitList) *orbitList {
	if rootBodies, ok := orbits[parentName]; ok {
		for _, body := range rootBodies {
			if parentOrbit, ok := fullOrbits.get(parentName); ok {
				orbit := orbit { name: body, parent: &parentOrbit }
				*fullOrbits = append(*fullOrbits, orbit)
			}

			addChildOrbitsOf(body, orbits, fullOrbits)
		}
	}

	return fullOrbits
}

func countAllOrbitTypes(orbits orbitList) int {
	total := 0
	for _, orbit := range orbits {
		orbitCounter(orbit, &total)
	}
	return total
}

func orbitCounter(currentOrbit orbit, count *int) {
	if currentOrbit.parent != nil {
		(*count)++
		orbitCounter(*currentOrbit.parent, count)
	}
}

// Solve solves both parts of the problem
func Solve() {
	day := 6

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

	preOrbits, err := readPreOrbits(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to read preOrbits")
		fmt.Fprintln(os.Stderr, err)
		return 
	}

	orbits := processOrbits(preOrbits)
	totalOrbits := countAllOrbitTypes(orbits)

	fmt.Printf("total orbits (direct & indirect) = %d\n", totalOrbits)
}
