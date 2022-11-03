package main

import "github.com/gofiber/fiber/v2"

func RegisterMainRoutes(app *fiber.App) {
	app.Get("/", home)
	app.Get("/posts/", getAllPosts)
	app.Get("/posts/:post", getPost)
	app.Get("/:page", getPage)
}
