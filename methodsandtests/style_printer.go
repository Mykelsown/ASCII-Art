package methodsandtests

import (
	"log"
	"strings"
)

func Printer(str string) string {
	// var all strings.Builder
	charFormatSlice := make([]string, len(str))
	strElSlice := strings.Split(str, "")
	for _, el := range strElSlice {
		// _, err := all.WriteString(GetFormattedStyle(el))
		// if err != nil {
		// 	log.Fatal("%w", err)
		// }

		charFormatSlice = append(charFormatSlice, GetFormattedStyle(el))
	}

	var fin strings.Builder
	unstructTwoDimSlice := make([][]string, 9)
	for i, formEl := range charFormatSlice {
		partOfFormEl := strings.Split(formEl, "\n")
		unstructTwoDimSlice[i] = partOfFormEl
		fin.WriteString(formEl)
	}
	
	return unstructTwoDimSlice[0][1]
}

func GetFormattedStyle(str string) string {
	bufferedStyleFile := FileHandler()
	line := 1
	chLine := 1
	var char strings.Builder
	for bufferedStyleFile.Scan() {
		for _, ch := range str {
			if ch >= ' ' || ch <= '~' {
				intCh := int(ch - ' ')
				chLine = intCh * 9
			}

			if line > chLine && line <= chLine+9 {
				char.WriteString(bufferedStyleFile.Text() + "\n")
				break
			}
		}
		line++
	}

	// for row := 0; row < 8; row++ {
	// 	for _, ch := range str {
	// 		char.WriteString(bufferedStyleFile.Text())
	// 	}

	// }

	if err := bufferedStyleFile.Err(); err != nil {
		log.Fatal(err)
	}
	return char.String()
}
