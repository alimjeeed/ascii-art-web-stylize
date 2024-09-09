package asciiart

// GetAsciiArt takes a slice of strings (outputLines) and combines them into a single string.
// Each element is separated by a newline, except for the last line.
func GetAsciiArt(outputLines []string) string {
	output := ""
	if !IsStringSliceEmpty(outputLines) {
		for i, line := range outputLines {
			if i == len(outputLines)-1 {
				output += line
			} else {
				output += line + "\n"
			}
		}
	}
	return output
}
