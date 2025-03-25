package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the number of arguments passed is less than 3
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not Enough Args")
		// Exit the program with a status code of -1 (indicates error)
		os.Exit(-1)
	}

	// Retrieve the first and second command-line arguments
	// `old` is the text to be replaced, `new` is the text to replace it with
	old, new := os.Args[1], os.Args[2]

	// Create a Scanner to read input from the standard input (stdin)
	scan := bufio.NewScanner(os.Stdin)

	// Loop through each line of input
	for scan.Scan() {
		// Split the current line of text using the `old` substring as a delimiter
		// This creates a slice containing parts of the string without `old`
		s := strings.Split(scan.Text(), old)
		// Join the slice back into a single string, inserting `new` between the parts
		t := strings.Join(s, new)
		// Print the modified line to standard output (stdout)
		fmt.Println(t)
	}
}