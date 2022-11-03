package main

import "github.com/gofiber/fiber/v2"

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getPage(c *fiber.Ctx) error {
	return c.SendString("page")
}

func getPost(c *fiber.Ctx) error {
	return c.SendString("post")
}

func getAllPosts(c *fiber.Ctx) error {
	return c.SendString("all post")
}
