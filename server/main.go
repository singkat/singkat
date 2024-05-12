package main

import (
	"os"
	"singkat/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	mux := fiber.New()
	mux.Use(cors.New())

	mux.Post("links", handler.PostLinks)
	mux.Get("/", handler.RedirectLink)
	err := mux.Listen(":3001")
	if err != nil {
		os.Exit(1)
	}
}
