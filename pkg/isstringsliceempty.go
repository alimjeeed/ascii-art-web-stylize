package asciiart

// IsStringSliceEmpty checks if all elements in the provided slice of strings are empty.
// Returns true if all elements are empty, otherwise returns false.
func IsStringSliceEmpty(slice []string) bool {
	for _, s := range slice {
		if s != "" {
			return false
		}
	}
	return true
}
