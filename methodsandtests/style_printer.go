package methodsandtests

import (
	"log"
	"strings"
)

func Printer(str string) (string, string, int) {
	bufferedStyleFile := FileHandler()
	line := 0
	var char strings.Builder
	for bufferedStyleFile.Scan() {
		if line > 9 || line < 18 {
			char.WriteString(bufferedStyleFile.Text())
			// break
		}
		line++
	}
	if err := bufferedStyleFile.Err(); err != nil {
		log.Fatal(err)
	}
	return str, char.String(), line
}
