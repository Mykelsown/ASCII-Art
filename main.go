package main

import (
	"asciiart/methodsandtests"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("The input format is wrong.\nCorrect format: go run main.go <text-argument>")
	}

	fmt.Print(methodsandtests.Printer(os.Args[1]))
}
