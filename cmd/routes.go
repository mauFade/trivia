package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/", handlers.CreateFact)
}
