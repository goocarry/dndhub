package main

import (
	"fmt"
	"log"
	"net/http"
)

// Config ...
type Config struct{}

// PORT ...
const PORT = "3004"

func main() {
	app := Config{}

	log.Println("Starting mail service on port", PORT)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
