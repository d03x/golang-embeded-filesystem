package main

import (
	"embed"
	"log"
	"net/http"

	"dadandev.com/golang-project/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed assets/*
var frontendAssets embed.FS

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "BACKEND AP",
	})

	app.Use(middleware.Logging)

	//get data
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("assets/index.html")
	})
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(frontendAssets),
	}))
	//get user data
	log.Fatal(app.Listen(":8005"))
}
