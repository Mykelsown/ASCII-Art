package methodsandtests



// BannerLoader holds the loaded banner data
type BannerLoader struct {
	charMap map[rune][]string
}

// PrintString converts a string to ASCII art using preloaded banner
func (b *BannerLoader) PrintString(str string) []string {
	if str == "" {
		return []string{}
	}

	// Initialize 8 lines for the output
	artLines := make([]string, 8)

	for _, char := range str {
		// Check if character is in valid ASCII range
		if char < ' ' || char > '~' {
			continue
		}

		// Get character lines from preloaded map
		charLines, exists := b.charMap[char]
		if !exists {
			// Use space if character not found
			charLines = b.charMap[' ']
		}

		// Append each line to the corresponding output line
		for i := 0; i < 8 && i < len(charLines); i++ {
			artLines[i] += charLines[i]
		}
	}

	return artLines
}

