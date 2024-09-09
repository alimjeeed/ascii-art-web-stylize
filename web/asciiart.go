package web

import (
	asciiart "asciiart/pkg"
	"net/http"
	"strings"
)

const (
	defaultBanner   = "standard"
	fileExtension   = ".txt"
	unixSplitter    = "\n"
	windowsSplitter = "\r\n"
)

// Define the data structure to be passed to the HTML template.
type PageData struct {
	Output string
}

// Handles HTTP POST requests to the "/ascii-art" URL.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Send a 500 Internal Server Error response to the client if the template does not exist.
	if tmpl == nil {
		http.Error(w, "Internal Server Error: Template not found.", http.StatusInternalServerError)
		return
	}

	// Check if the request method is POST.
	if r.Method == http.MethodPost {
		inputString := r.FormValue("inputString")
		banner := r.FormValue("banner")

		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			banner = defaultBanner
		}

		// Ensure that the input string is not empty.
		if inputString == "" {
			// Send a 400 Bad Request response to the client.
			http.Error(w, "Bad Request: Input string cannot be empty. Please provide a valid input.", http.StatusBadRequest)
			return
		}

		// Ensure that the input string contains only ASCII characters.
		for _, v := range inputString {
			if !(v >= 32 && v <= 126) {
				// Send a 400 Bad Request response to the client.
				http.Error(w, "Bad Request: Input must contain only printable ASCII characters. Please provide a valid input.", http.StatusBadRequest)
				return
			}
		}

		// To solve the thinkertoy unix/windows problem.
		splitter := unixSplitter
		if banner == "thinkertoy" {
			splitter = windowsSplitter
		}

		// Validate banner file existence, then read and split its content into lines.
		bannerFileName := banner + fileExtension
		asciiart.ValidateBannerFile(bannerFileName, w, r)
		asciiArtLines := strings.Split(asciiart.ReadFileContent(bannerFileName), splitter)

		// Convert input string to ASCII values and generate/display ASCII art.
		asciiValues := asciiart.StringToAscii(inputString)
		output := asciiart.GenerateAsciiArt(asciiValues, asciiArtLines)

		// Create an instance of PageData, which holds the data to be passed to the HTML template.
		data := PageData{
			Output: output,
		}

		// Execute the template, passing in the data to be displayed.
		tmpl.Execute(w, data)

		// Added redirection to the home page for GET requests to prevent displaying a blank page.
	} else if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
