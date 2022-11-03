package main

import (
	"github.com/gofiber/fiber/v2"
	"icecreamcode.org/bubblegum/bubblegum_api/v1"
)

func main() {
	// get data
	bubblegum_api.LoadData()

	// create server
	app := fiber.New()
	app.Static("/", "./public")

	RegisterMainRoutes(app)
	bubblegum_api.RegisterRoutes(app)

	app.Listen(":3000")
}
