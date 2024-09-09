package web

import (
	"net/http"
	"os"
	"text/template"
)

// Checks if the "templates/index.html" file exists.
func CheckTemplateExists() bool {
	const templatePath = "templates/index.html"
	_, err := os.Stat(templatePath)
	return err == nil
}

var tmpl *template.Template

func init() {
	// Check if the template exists before parsing it.
	if CheckTemplateExists() {
		tmpl = template.Must(template.ParseFiles("templates/index.html"))
	}
}

// Handles HTTP GET requests to the root URL ("/").
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// If the url is wrong, send a 404 Not Found response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Send a 500 Internal Server Error response to the client if the template does not exist.
	if tmpl == nil {
		http.Error(w, "Internal Server Error: Template not found.", http.StatusInternalServerError)
		return
	}

	// Check if the request method is GET.
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	}
}
