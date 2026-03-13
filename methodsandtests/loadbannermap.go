package methodsandtests

import (
	"fmt"
)

func LoadBannerMap() (map[rune][]string, error) {
	bufferedStyleFile := FileHandler()
	if bufferedStyleFile == nil {
		return nil, fmt.Errorf("failed to open banner file")
	}

	bannerMap := make(map[rune][]string)
	currentChar := rune(32) // Start with space (ASCII 32)
	charLines := []string{}

	for bufferedStyleFile.Scan() {
		line := bufferedStyleFile.Text()

		// Empty line means end of current character (separator)
		if line == "" {
			if len(charLines) > 0 {
				// Store the collected lines for this character
				bannerMap[currentChar] = make([]string, len(charLines))
				copy(bannerMap[currentChar], charLines)

				// Reset for next character
				charLines = []string{}
				currentChar++
			}
		} else {
			// Add line to current character
			charLines = append(charLines, line)
		}
	}

	// Don't forget the last character if file doesn't end with empty line
	if len(charLines) > 0 {
		bannerMap[currentChar] = charLines
	}

	if err := bufferedStyleFile.Err(); err != nil {
		return nil, err
	}

	return bannerMap, nil
}
