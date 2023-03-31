package main

import (
	"errors"
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
	if len(envFilesDir) > 0 {
		filesDir = envFilesDir
	}

	log.Println("env: ", filesDir)

	if _, err := os.Stat(filesDir); errors.Is(err, os.ErrNotExist) {
		log.Println("FILE_DIR NOT EXISTS")
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
