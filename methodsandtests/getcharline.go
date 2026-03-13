package methodsandtests

// GetCharLine gets a specific line (0-7) for a character
func (b *BannerLoader) GetCharLine(char rune, lineNum int) string {
	charLines, exists := b.charMap[char]
	if !exists || lineNum < 0 || lineNum >= len(charLines) {
		return ""
	}
	return charLines[lineNum]
}
