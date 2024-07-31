package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/trivia/db"
)

func main() {
	db.OpenDatabase()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("VAMOSSSSSS")
	})

	app.Listen(":3000")
}
