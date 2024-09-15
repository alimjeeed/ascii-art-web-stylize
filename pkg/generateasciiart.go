package asciiart

const (
	asciiArtHeight      = 8
	asciiOffset         = 32
	linesPerCharacter   = 9
	asciiNewline        = 10
	asciiCarriageReturn = 13
)

// GenerateAsciiArt converts a list of ASCII values into an ASCII art representation.
// It processes each character, handles newlines, and constructs the final output string.
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

		// Check if the current character starts a new line.
		if i < len(asciiValues)-1 && asciiValues[i] == asciiNewline || asciiValues[i] == asciiCarriageReturn {
			
			// Append the ASCII art created so far and add a newline.
			output += GetAsciiArt(outputLines) + "\n"

			// Reset the outputLines slice to start a new line of ASCII art.
			outputLines = make([]string, asciiArtHeight)

			// Skip the newline character in the input.
			skipNext = true
			continue
		}

		// Calculate the starting index in the banner file for the current character.
		startIndex := (asciiValues[i]-asciiOffset)*linesPerCharacter + 1

		// Build the ASCII art for the current character.
		for j := 0; j < asciiArtHeight; j++ {
			outputLines[j] += asciiArtLines[startIndex+j]
		}

		// Check if the last character in the input string is a newline.
		if len(asciiValues) > 1 && asciiValues[len(asciiValues)-1] == asciiNewline || asciiValues[len(asciiValues)-1] == asciiCarriageReturn {
			lastCharIsNewLine = true
		}

		// If this is the last character, append the remaining ASCII art.
		if i == len(asciiValues)-1 {
			output += GetAsciiArt(outputLines)
		}
	}

	// Add the newline at the end if required.
	if lastCharIsNewLine {
		output += "\n"
	}

	return output
}
