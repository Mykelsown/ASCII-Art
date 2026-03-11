package methodsandtests

import (
	"log"
	"strings"
)

func Printer(str string) (string, string, int, int) {
	bufferedStyleFile := FileHandler()
	line := 1
	chLine := 0
	var char strings.Builder
	for bufferedStyleFile.Scan() {
		for _, ch := range str {
			charStartingLine := make([]int, len(str))
			if ch >= ' ' || ch <= '~' {
				intCh := int(ch - ' ')
				chLine = intCh * 9 + 1
				charStartingLine = append(charStartingLine, chLine)
			}
		}
		char.WriteString(bufferedStyleFile.Text() + "\n")
		if line == 18 {
			break
		}
		line++
	}
	if err := bufferedStyleFile.Err(); err != nil {
		log.Fatal(err)
	}
	return str, char.String(), line, chLine
}
