package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PostLinks(c *fiber.Ctx) error {
	in := fiber.Map{}
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    "https://singkat.com/?to=" + in["url"].(string),
	})
}
