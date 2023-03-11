package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	logger := log.New(os.Stdout, "collector: ", 0)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("collector page is requested")
		render(w, "test.page.gohtml")
	})

	logger.Println("Starting collector server...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	partials := []string{
		"./templates/base.layout.gohtml",
		"./templates/header.partial.gohtml",
		"./templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

	for _, t := range partials {
		templateSlice = append(templateSlice, t)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
