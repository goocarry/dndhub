package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// PORT ...
	PORT = "3006"
)

// Config ...
type Config struct{}

func main() {
	app := Config{}

	// start web server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}
	log.Println("Starting service on port", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
