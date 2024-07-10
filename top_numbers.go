package main

import (
	"bufio"      // implements buffered I/O. It provides buffering for reading and writing operations, which can improve performance when handling large amounts of data.
	"errors"     // implements functions to manipulate errors. It provides utilities for creating and working with errors.
	"flag"       // implements command-line flag parsing. It allows parsing and handling of command-line arguments passed to the program.
	"fmt"        // implements formatted I/O with functions analogous to C's printf and scanf. It provides methods for printing formatted output and reading formatted input.
	"os"         // provides a platform-independent interface to operating system functionality. It allows accessing operating system functionalities like file operations, environment variables, etc.
	"sort"       // provides primitives for sorting slices and user-defined collections. It enables sorting of data collections in various orders.
	"strconv"    // implements conversions to and from string representations of basic data types. It provides functions for parsing strings into numeric types and vice versa.
)

const MaxTopResults = 30000000 // Maximum number of top results allowed

func main() {
	// Define command line arguments
	n := flag.Int("n", 0, "Number of top results to output")
	inputFile := flag.String("input-file", "", "Path to the input file")
	outputFile := flag.String("output-file", "", "Path to the output file")
	flag.Parse() // Parse command line arguments

	// Validate the N argument
	if *n <= 0 {
		fmt.Println("ERROR: Number of top results must be bigger than 0.")
		os.Exit(1)
	}
	if *n > MaxTopResults {
		fmt.Println("ERROR: The maximum number of top results must be less or equal than 30000000.")
		os.Exit(1)
	}

	// Check if input file exists and is readable
	if _, err := os.Stat(*inputFile); errors.Is(err, os.ErrNotExist) {
		fmt.Println("ERROR: input file does not exist.")
		os.Exit(1)
	}

	file, err := os.Open(*inputFile) // Open input file for reading
	if err != nil {
		fmt.Println("ERROR: input file is not readable.")
		os.Exit(1)
	}
	defer file.Close() // Ensure file is closed when function exits

	// Read numbers from the file
	var numbers []uint64 // Slice to hold valid numbers
	scanner := bufio.NewScanner(file) // Scanner for reading file line by line
	lineNumber := 0 // Counter for line number

	for scanner.Scan() { // Loop through each line in the file
		lineNumber++
		line := scanner.Text() // Get the current line
		num, err := strconv.ParseUint(line, 10, 64) // Convert line to uint64
		if err != nil { // Handle invalid lines
			fmt.Printf("WARN: Invalid line %d\n", lineNumber)
			continue // Skip to next line
		}
		numbers = append(numbers, num) // Append valid number to slice
	}

	if err := scanner.Err(); err != nil { // Check for scanner errors
		fmt.Println("ERROR: Failed to read the input file.")
		os.Exit(1)
	}

	// Sort numbers in descending order and select top N
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j] // Sort in descending order
	})
	if len(numbers) > *n {
		numbers = numbers[:*n] // Take only top N numbers
	}

	// Write the top N numbers to the output file
	outFile, err := os.Create(*outputFile) // Create output file
	if err != nil {
		fmt.Println("ERROR: Failed to create the output file.")
		os.Exit(1)
	}
	defer outFile.Close() // Ensure file is closed when function exits

	writer := bufio.NewWriter(outFile) // Writer for writing to file
	for _, num := range numbers {
		_, err := fmt.Fprintln(writer, num) // Write each number to file
		if err != nil {
			fmt.Println("ERROR: Failed to write to the output file.")
			os.Exit(1)
		}
	}
	writer.Flush() // Flush buffer to ensure all data is written
}
