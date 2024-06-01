package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// reads user input until \n by default
	scanner.Scan()

	// Holds the string that was scanned
	text := scanner.Text()

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	// separate the strings into array
	input := strings.Fields(text)
	// check if input is more then 3 number
	if len(input) != 3 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	// convert every input into int
	width, err := strconv.Atoi(input[0])
	// fail to convert or check if faulty inputs
	if err != nil || width < 1 || width > 50000 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	length, err := strconv.Atoi(input[1])
	if err != nil || length < 1 || length > 50000 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	count, err := strconv.Atoi(input[2])
	if err != nil || count < 1 || count > 50000 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	estate := createEstate(width, length)

	for i := 0; i < count; i++ {
		// reads user input until \n by default
		scanner.Scan()

		// Holds the string that was scanned
		text := scanner.Text()

		// handle error
		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		// separate the strings into array
		input := strings.Fields(text)
		// check if input is more then 3 number
		if len(input) != 3 {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}

		// convert every input into int
		x, err := strconv.Atoi(input[0])
		// fail to convert or faultyy inputs
		if err != nil || x < 1 || x > width {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		y, err := strconv.Atoi(input[1])
		if err != nil || y < 1 || y > length {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		height, err := strconv.Atoi(input[2])
		if err != nil || height < 0 || height > 30 {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		// put tree height on estate(x,y)
		estate[length-y][x-1] = height
	}
	fmt.Println(estate)

	// fail := false
	// if fail {
	// 	// On input faulty, print "FAIL" to stderr and exit with status 1
	// 	fmt.Fprintln(os.Stderr, "FAIL")
	// 	os.Exit(1)
	// }
	// // On input correct, print the result to stdout and exit with status 0
	// output := 8
	// fmt.Printf("%d\n", output)

	output := countDroneDistance(estate)
	fmt.Printf("%d\n", output)
}

func createEstate(width int, length int) [][]int {
	// initiate a empty 2d array
	estate := [][]int{}

	for i := 0; i < length; i++ {
		// we fill each row an array of 0s with the width size
		l := make([]int, width)
		// insert it into estate
		estate = append(estate, l)
	}

	return estate
}

func countDroneDistance(estate [][]int) int {
	// distance per plot
	distPerPlot := 10
	// initiate drone position
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
					diff := abs(currDroneHeight - (estate[row][column] + 1))
					totalDistance = totalDistance + diff
					currDroneHeight = estate[row][column] + 1
				}

			}
		} else if column == len(estate[row]) {
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
					diff := abs(currDroneHeight - (estate[row][column] + 1))
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
			diff := abs(currDroneHeight - (estate[row][column] + 1))
			totalDistance = totalDistance + diff
		}
	}

	// at the last plot drone is goinng to go to ground
	totalDistance = totalDistance + currDroneHeight

	return totalDistance
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
