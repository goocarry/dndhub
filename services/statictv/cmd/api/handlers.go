package main

import (
	"net/http"
)

// StreamStatic ...
func (app *Config) StreamStatic() http.Handler {

	fs := http.FileServer(http.Dir("../../static"))

	return fs
}
