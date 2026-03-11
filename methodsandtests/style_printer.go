package methodsandtests

import (
	"log"
	"strings"
)

func Printer(str string) (string, int, int) {
	bufferedStyleFile := FileHandler()
	line := 1
	chLine := 0
	var char strings.Builder
	for bufferedStyleFile.Scan() {
		for _, ch := range str {
			charStartingLine := make([]int, len(str))
			if ch >= ' ' || ch <= '~' {
				intCh := int(ch - ' ')
				chLine = intCh*9 
				charStartingLine = append(charStartingLine, chLine)
			}
			if line >= int(chLine) && line <= int(chLine)+9 {
				char.WriteString(bufferedStyleFile.Text() + "\n")
				break
			}
		}
		line++
	}
	if err := bufferedStyleFile.Err(); err != nil {
		log.Fatal(err)
	}
	return  char.String(), line, chLine
}
