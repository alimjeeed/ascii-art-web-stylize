package asciiart

const (
	asciiArtHeight    = 8
	asciiOffset       = 32
	linesPerCharacter = 9
)

func GenerateAsciiArt(asciiValues []int, asciiArtLines []string) string {
	outputLines := make([]string, asciiArtHeight)
	skipNext := false
	lastCharIsNewLine := false

	output := ""

	// This Loop iterates through each character in the input string.
	for i := 0; i < len(asciiValues); i++ {
		if skipNext {
			skipNext = false
			continue
		}

		// If the current and next character form "\n", insert a new line.
		if i < len(asciiValues)-1 && asciiValues[i] == '\\' && asciiValues[i+1] == 'n' {
			// Add the ASCII art created so far, followed by a newline character.
			output += GetAsciiArt(outputLines) + "\n"

			// Reset the outputLines slice for the next line of ASCII art.
			outputLines = make([]string, asciiArtHeight)

			// Skip "\n"
			skipNext = true
			continue
		}

		// Calculate the starting index in the banner file for the current character.
		startIndex := (asciiValues[i]-asciiOffset)*linesPerCharacter + 1

		// Build the ASCII art for the current character.
		for j := 0; j < asciiArtHeight; j++ {
			outputLines[j] += asciiArtLines[startIndex+j]
		}

		// Check if the last two characters in the input string form a newline.
		if len(asciiValues) > 1 && asciiValues[len(asciiValues)-1] == 'n' && asciiValues[len(asciiValues)-2] == '\\' {
			lastCharIsNewLine = true
		}

		// If this is the last character, add the remaining ASCII art.
		if i == len(asciiValues)-1 {
			output += GetAsciiArt(outputLines)
		}
	}

	// Add the new line at the end if required.
	if lastCharIsNewLine {
		output += "\n"
	}

	return output
}
