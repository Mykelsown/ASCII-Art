package main

import (
	"asciiart/methodsandtests"
	"fmt"
	"os"
)

func main() {
	// Check arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <string>")
		return
	}

	input := os.Args[1]

	// Handle empty string
	if input == "" {
		return
	}

	// Handle single newline
	if input == "\\n" {
		fmt.Println()
		return
	}

	// STEP 1: Load the banner file once
	banner, err := methodsandtests.NewBannerLoader()
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	// STEP 2: Convert string to ASCII art
	artLines := banner.PrintString(input)

	// STEP 3: Print each line
	for _, line := range artLines {
		fmt.Println(line)
	}
}