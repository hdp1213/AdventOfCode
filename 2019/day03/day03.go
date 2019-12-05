package day03

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/hdp1213/AdventOfCode/2019/utils"
)

type point struct {
	x int
	y int
}

type lineSegment struct {
	start point
	end point
}

func readWirePaths(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result, scanner.Err()
}

func getWirePathElements(wirePath string) ([]string, error) {
	stringReader := strings.NewReader(wirePath)
	scanner := bufio.NewScanner(stringReader)
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

	var pathElements []string
	for scanner.Scan() {
		elem := scanner.Text()
		pathElements = append(pathElements, elem)
	}

	return pathElements, scanner.Err()
}

func getCoords(wireElements []string) (points []point, err error) {
	// As a reminder, x is R/L, y is U/D. +ve x -> R, +ve y -> U
	newPoint := point {x:0,y:0}
	points = make([]point, len(wireElements) + 1)
	points[0] = newPoint

	for i, element := range wireElements {
		direction := element[0]
		lengthString := element[1:len(element)]

		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return points, err
		}

		if direction == 'R' {
			newPoint.x = points[i].x + length
		}
		if direction == 'L' {
			newPoint.x = points[i].x - length
		}

		if direction == 'U' {
			newPoint.y = points[i].y + length
		}
		if direction == 'D' {
			newPoint.y = points[i].y - length
		}

		points[i + 1] = newPoint
	}

	return
}

func doesIntersect(seg1 lineSegment, seg2 lineSegment) (intersection bool) {
	intersection = false
	return
}

// Solve solves both parts of the problem
func Solve() {
	day := 3
	inputFile, err := utils.GetInput(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bad things happened")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "more bad things happened")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer file.Close()

	paths, err := readWirePaths(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading wire paths failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	firstWire, err := getWirePathElements(paths[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading wire paths failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	points, err := getCoords(firstWire)
	if err != nil {
		fmt.Fprintln(os.Stderr, "calculating wire coords failed")
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(points)
}
