package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// PORT ...
const PORT = "3007"

func main() {
	logger := log.New(os.Stdout, "display service: ", 0)
	logger.Println("starting display service...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("display page is requested")
		render(w, "test.page.gohtml", logger)
	})

	logger.Printf("Starting display server at port %s...", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
	if err != nil {
		logger.Panic(err)
	}
}


//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string, logger *log.Logger) {
	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	for _, t := range partials {
		templateSlice = append(templateSlice, t)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Panic(err)
		return
	}

	var data struct {
		StreamURL string
	}

	data.StreamURL = os.Getenv("StreamURL")

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Panic(err)
		return
	}
}
