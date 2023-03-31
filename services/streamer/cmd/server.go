package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// PORT ...
const PORT = 3005

// Run ...
func Run() error {
	log.Printf("Starting streamer server on port: %d", PORT)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New())
	app.Use(cors.New())

	return app.Listen(fmt.Sprintf(":%d", PORT))
}
