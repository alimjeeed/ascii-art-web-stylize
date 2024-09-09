package asciiart

import (
	"net/http"
	"os"
)

// ValidateBannerFile checks if the specified banner file exists.
// If the file does not exist, it sends a 404 Not Found response to the client.
func ValidateBannerFile(banner string, w http.ResponseWriter, r *http.Request) {
	bannerPath := "banners/" + banner
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
}
