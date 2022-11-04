package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"icecreamcode.org/bubblegum/bubblegum_api/v1"
)

func main() {
	// get data
	bubblegum_api.LoadData()

	engine := html.New("./templates", ".html")

	// create Fiber
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Static("/", "./public")

	RegisterMainRoutes(app)
	bubblegum_api.RegisterRoutesV1(app)

	app.Listen(":3000")
}
