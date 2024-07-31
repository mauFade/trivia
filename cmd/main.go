package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/trivia/db"
)

func main() {
	db.OpenDatabase()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
