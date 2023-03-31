package main

import "log"

func main() {
	if err := Run(); err != nil {
		log.Panic("error with streamer server: %w", err)
	}
}
