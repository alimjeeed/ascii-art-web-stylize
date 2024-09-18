package asciiart

// StringToAscii converts each character in the input string to its corresponding ASCII value.
// And returns a slice of these integer ASCII values.
func StringToAscii(s string) []int {
	asciiValues := []int{}
	for _, char := range s {
		asciiValues = append(asciiValues, int(char))
	}
	return asciiValues
}
