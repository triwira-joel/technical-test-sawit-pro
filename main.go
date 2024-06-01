package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// This is just a placeholder to show how to print to stderr and return a non-zero exit code
	// Please replace the whole content of this function with your solution.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter width, length, count: ")
	// reads user input until \n by default
	scanner.Scan()
	// Holds the string that was scanned
	text := scanner.Text()
	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	words := strings.Fields(text)
	if len(words) != 3 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	width, err := strconv.Atoi(words[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	length, err := strconv.Atoi(words[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	count, err := strconv.Atoi(words[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	estate := createEstate(width, length)
	fmt.Println(estate)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter Tree #%d  x y height: \n", i+1)
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()

		// handle error
		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		words := strings.Fields(text)
		if len(words) != 3 {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		x, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		y, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		height, err := strconv.Atoi(words[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}
		// put tree height on estate x y
		estate[y-2+length][x-1] = height
	}

	fmt.Println(estate)

	fail := false
	if fail {
		// On input faulty, print "FAIL" to stderr and exit with status 1
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	// On input correct, print the result to stdout and exit with status 0
	output := 8
	fmt.Printf("%d\n", output)
}

func createEstate(width int, length int) [][]int {
	estate := [][]int{}
	for i := 0; i < width; i++ {
		l := make([]int, length)
		estate = append(estate, l)
	}
	return estate
}
