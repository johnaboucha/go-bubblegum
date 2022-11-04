package main

import (
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	cats := pullCategories(c.BaseURL())

	return c.Render("home", fiber.Map{
		"Title": "Hello, World!",
		"Cats":  cats,
	})
}
