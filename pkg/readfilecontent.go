package asciiart

import (
	"os"
)

// ReadFileContent reads the contents of a specified banner file and returns it as a string.
// If an error occurs (e.g., file not found), the function will panic.
func ReadFileContent(banner string) string {
	bannerPath := "banners/" + banner
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
