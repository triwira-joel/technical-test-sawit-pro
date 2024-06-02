package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	h "github.com/triwira-joel/technical-test-sawit-pro/helper"
)

func main() {
	estate, success := ReadInput()
	if !success {
		// On input faulty, print "FAIL" to stderr and exit with status 1
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	// On input correct, print the result to stdout and exit with status 0
	output := CountDroneDistance(*estate)
	fmt.Printf("%d\n", output)
}

func ReadInput() (*[][]int, bool) {
	scanner := bufio.NewScanner(os.Stdin)

	// reads user input until \n by default
	scanner.Scan()

	// Holds the string that was scanned
	text := scanner.Text()

	// separate the strings into array
	input := strings.Fields(text)
	// check if input is more/less then 3 number
	if len(input) != 3 {
		return nil, false
	}

	// convert every input into int
	width, err := strconv.Atoi(input[0])
	// fail to convert or check if faulty inputs
	if err != nil || width < 1 || width > 50000 {
		return nil, false
	}
	length, err := strconv.Atoi(input[1])
	if err != nil || length < 1 || length > 50000 {
		return nil, false
	}
	count, err := strconv.Atoi(input[2])
	if err != nil || count < 1 || count > width*length {
		return nil, false
	}

	estate := h.CreateEstate(width, length)

	for i := 0; i < count; i++ {
		// reads user input until \n by default
		scanner.Scan()

		// Holds the string that was scanned
		text := scanner.Text()

		// separate the strings into array
		input := strings.Fields(text)
		// check if input is more/less then 3 number
		if len(input) != 3 {
			return nil, false
		}

		// convert every input into int
		x, err := strconv.Atoi(input[0])
		// fail to convert or faultyy inputs
		if err != nil || x < 1 || x > width {
			return nil, false
		}
		y, err := strconv.Atoi(input[1])
		if err != nil || y < 1 || y > length {
			return nil, false
		}
		height, err := strconv.Atoi(input[2])
		if err != nil || height < 0 || height > 30 {
			return nil, false
		}
		// put tree height on estate(x,y)
		estate[length-y][x-1] = height
	}

	return &estate, true
}

func CountDroneDistance(estate [][]int) int {
	distPerPlot := 10
	// initiate drone position from southwest
	row, column := len(estate)-1, 0
	totalDistance := 0

	// initiate drone distance from southwest corner
	if estate[row][column] != 0 {
		totalDistance = totalDistance + estate[row][column] + 1 // if there is a tree
	} else {
		totalDistance++ // if no tree
	}
	// initiate a variable to track drone height
	currDroneHeight := totalDistance

	// looping each row to north
	for row >= 0 {
		if column == 0 { // if its from west side
			// looping to go to east
			for column < len(estate[row]) {
				// if already at the end, break
				if column == len(estate[row])-1 {
					break
				}
				// we go right
				column++
				totalDistance = totalDistance + distPerPlot

				// check if current drone height is the same as next plot
				if currDroneHeight != estate[row][column]+1 {
					// we travel upward to ground/tree's height
					diff := h.Abs(currDroneHeight - (estate[row][column] + 1))
					totalDistance = totalDistance + diff
					currDroneHeight = estate[row][column] + 1
				}

			}
		} else {
			// looping to go to west
			for column >= 0 {
				// if already at the end, break
				if column == 0 {
					break
				}
				// we go left
				column--
				totalDistance = totalDistance + distPerPlot

				// check if current drone height is the same as next plot
				if currDroneHeight != estate[row][column]+1 {
					// we travel upward/downward to ground/tree's height
					diff := h.Abs(currDroneHeight - (estate[row][column] + 1))
					totalDistance = totalDistance + diff
					currDroneHeight = estate[row][column] + 1
				}
			}
		}
		// condition where we cant go north
		if row == 0 {
			break
		}
		// we go north
		row--
		totalDistance = totalDistance + distPerPlot
		if currDroneHeight != estate[row][column]+1 {
			// we travel upward to ground/tree's height
			diff := h.Abs(currDroneHeight - (estate[row][column] + 1))
			totalDistance = totalDistance + diff
			currDroneHeight = estate[row][column] + 1
		}
	}

	// at the last plot drone is goinng to go to ground
	totalDistance = totalDistance + currDroneHeight

	return totalDistance
}
