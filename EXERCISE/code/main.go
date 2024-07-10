package main

import (
	"bufio" // Package bufio implements buffered I/O.
	"flag"  // Package flag implements command-line flag parsing.
	"fmt"   // Package fmt implements formatted I/O.
	"math"  // Package math provides basic constants and mathematical functions.
	"os"    // Package os provides a platform-independent interface to operating system functionality.
	"sort"  // Package sort provides primitives for sorting slices and user-defined collections.
	"strconv" // Package strconv implements conversions to and from string representations of basic data types.
)

// Constants to define maximum allowable results and valid integer value.
const (
	maxTopResults = 30000000
	maxValidValue = math.MaxInt64 // 9223372036854775807
)

func main() {
	// Define command-line flags.
	n := flag.Int("n", 0, "Number of top results")
	inputFile := flag.String("input-file", "", "Path to input file")
	outputFile := flag.String("output-file", "", "Path to output file")
	flag.Parse() // Parse the command-line flags.

	// Validate the number of top results.
	if *n <= 0 {
		fmt.Println("ERROR: Number of top results must be bigger than 0.")
		os.Exit(1)
	}

	if *n > maxTopResults {
		fmt.Printf("ERROR: The maximum number of top results must be less or equal than %d.\n", maxTopResults)
		os.Exit(1)
	}

	// Open the input file.
	file, err := os.Open(*inputFile)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: input file does not exist.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("ERROR: input file is not readable.")
		os.Exit(1)
	}
	defer file.Close() // Ensure the file is closed after the function finishes.

	// Initialize a scanner to read the input file line by line.
	scanner := bufio.NewScanner(file)
	// Create a slice to store the numbers, initially with capacity of n.
	numbers := make([]int64, 0, *n)
	lineNum := 0

	// Read the file line by line.
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		// Parse the line to an unsigned 64-bit integer.
		num, err := strconv.ParseUint(line, 10, 64)
		if err != nil || num > uint64(maxValidValue) {
			fmt.Printf("WARN: Invalid line %d.\n", lineNum)
			continue // Skip invalid lines.
		}
		numbers = append(numbers, int64(num)) // Add the valid number to the slice.
		// If the slice exceeds n elements, sort and trim it.
		if len(numbers) > *n {
			sort.Slice(numbers, func(i, j int) bool {
				return numbers[i] > numbers[j]
			})
			numbers = numbers[:*n] // Keep only the top n elements.
		}
	}

	// Check for scanning errors.
	if err := scanner.Err(); err != nil {
		fmt.Printf("ERROR: Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Final sort to ensure the top n elements are in order.
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// Create the output file.
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("ERROR: Unable to create output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close() // Ensure the file is closed after the function finishes.

	// Initialize a writer to write to the output file.
	writer := bufio.NewWriter(outFile)
	// Write the sorted numbers to the output file.
	for _, num := range numbers {
		_, err := fmt.Fprintln(writer, num)
		if err != nil {
			fmt.Printf("ERROR: Unable to write to output file: %v\n", err)
			os.Exit(1)
		}
	}

	// Flush the writer to ensure all data is written to the file.
	if err := writer.Flush(); err != nil {
		fmt.Printf("ERROR: Unable to flush output file: %v\n", err)
		os.Exit(1)
	}
}
