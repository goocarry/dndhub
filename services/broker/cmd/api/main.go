package main

import (
	"fmt"
	"log"
	"net/http"
)

// PORT ...
const PORT = "3000"

// Config ..
type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s", PORT)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
