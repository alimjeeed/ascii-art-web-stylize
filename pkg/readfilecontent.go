package asciiart

import (
	"fmt"
	"os"
)

// ReadFileContent reads the contents of a specified banner file and returns it as a string.
// If an error occurs (e.g., file not found), the function will panic.
func ReadFileContent(banner string) (string, error) {
	bannerPath := "banners/" + banner
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		return "", fmt.Errorf("banner file not found: %v", err)
	}
	return string(data), nil
}
