package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// PORT ...
	PORT = "3006"
	// DEFAULTDIR ...
	DEFAULTDIR = "../../static"
)

var filesDir = DEFAULTDIR

// Config ...
type Config struct{}

func main() {
	envFilesDir := os.Getenv("FILES_DIR")
	if len(filesDir) > 0 {
		filesDir = envFilesDir
	}

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
