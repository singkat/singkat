package handler

import (
	"fmt"
	"net/http"
	"singkat/utils"

	"github.com/gofiber/fiber/v2"
)

type UserRequest struct {
	Links string `json:"links"`
}

var LinkStore map[string]string

func PostLinks(c *fiber.Ctx) error {
	LinkStore = make(map[string]string)
	in := UserRequest{}
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusUnprocessableEntity).SendString(err.Error())
	}
	shortKey := utils.GenereateShortKey()

	fmt.Println("JADI")
	LinkStore[shortKey] = in.Links

	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    "localhost:3001/?to=" + shortKey,
	})
}

func RedirectLink(c *fiber.Ctx) error {
	shortKey := c.Query("to")
	
	if shortKey == "" {
		return c.Status(http.StatusInternalServerError).SendString("missing shortkey")
	}

	orginalLink, isFound := LinkStore[shortKey]

	if !isFound {
		return c.Status(http.StatusNotFound).SendString("Original Link not found, Please make new shortenLink again")
	}

	return c.Redirect(orginalLink, fiber.StatusPermanentRedirect)
}
