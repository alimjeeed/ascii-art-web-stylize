package asciiart

import (
	"fmt"
	"os"
)

// ReadBannerFile reads the contents of a specified banner file and returns it as a string.
func ReadBannerFile(bannerFileName string) (string, error) {
	bannerPath := "banners/" + bannerFileName
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		return "", fmt.Errorf("banner file not found: %v", err)
	}
	return string(data), nil
}
