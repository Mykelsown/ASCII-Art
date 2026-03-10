package methodsandtests

import (
	"bufio"
	"log"
	"os"
)

func FileHandler() *bufio.Scanner {
	styleFile, Err := os.Open("standard.txt")
	if Err != nil {
		log.Fatal(Err)
	}
	// defer styleFile.Close()

	bufferedStyleFile := bufio.NewScanner(styleFile)

	return bufferedStyleFile
}
