package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// PORT ...
const PORT = "3000"

func main() {

	logger := log.New(os.Stdout, "collector: ", 0)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("collector page is requested")
		render(w, "test.page.gohtml")
	})

	logger.Printf("Starting collector server at port %s...", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
	if err != nil {
		logger.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {
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
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
