package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/trivia/db"
	"github.com/mauFade/trivia/models"
)

func Home(c *fiber.Ctx) error {
	facts := []models.Fact{}
	db.Database.DB.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db.Database.DB.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}
