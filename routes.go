package main

import "github.com/gofiber/fiber/v2"

func RegisterMainRoutes(app *fiber.App) {
	app.Get("/", home)
}
