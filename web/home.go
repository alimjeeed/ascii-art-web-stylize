package web

import (
	"net/http"
	"os"
	"text/template"
)

const templatePath = "templates/index.html"

func CheckTemplateExists() bool {
	_, err := os.Stat(templatePath)
	return err == nil
}

var tmpl *template.Template

func init() {
	if CheckTemplateExists() {
		tmpl = template.Must(template.ParseFiles(templatePath))
	}
}

// Handles HTTP GET requests to the root URL ("/").
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// If the url is wrong, send a 404 Not Found response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if tmpl == nil {
		http.Error(w, "Internal Server Error: Template not found.", http.StatusInternalServerError)
		return
	}

	// Check if the request method is GET.
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	}
}
