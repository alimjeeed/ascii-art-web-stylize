package web

import (
	asciiart "asciiart/pkg"
	"net/http"
	"strings"
)

const (
	defaultBanner       = "standard"
	fileExtension       = ".txt"
	unixSplitter        = "\n"
	windowsSplitter     = "\r\n"
	asciiNewline        = 10
	asciiCarriageReturn = 13
)

// Define the data struct to be passed to the HTML template.
type PageData struct {
	Output string
}

// Handles HTTP POST requests to the "/ascii-art" URL.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	
	if tmpl == nil {
		http.Error(w, "Internal Server Error: Template not found.", http.StatusInternalServerError)
		return
	}

	// Check if the request method is POST.
	if r.Method == http.MethodPost {
		inputString := r.FormValue("inputString")
		banner := r.FormValue("banner")

		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			http.Error(w, "Bad Request: Invalid banner value.", http.StatusBadRequest)
			return
		}

		if inputString == "" {
			http.Error(w, "Bad Request: Input string cannot be empty. Please provide a valid input.", http.StatusBadRequest)
			return
		}

		// Ensure that the input string contains only ASCII characters.
		for _, v := range inputString {
			if !(v >= 32 && v <= 126) && v != asciiNewline && v != asciiCarriageReturn {
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
		bannerFileContent, err := asciiart.ReadFileContent(bannerFileName)
		if err != nil {
			http.Error(w, "Internal Server Error: " + err.Error(), http.StatusInternalServerError)
			return
		} 
		asciiArtLines := strings.Split(bannerFileContent, splitter)

		asciiValues := asciiart.StringToAscii(inputString)
		output := asciiart.GenerateAsciiArt(asciiValues, asciiArtLines)
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
