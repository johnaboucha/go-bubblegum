package bubblegum_api

import "github.com/gofiber/fiber/v2"

func apiHome(c *fiber.Ctx) error {
	return c.SendString("API Documentation")
}

func getResources(c *fiber.Ctx) error {
	return c.SendString("resources")
}

func getCategories(c *fiber.Ctx) error {
	keys := []string{}
	for k := range categories {
		keys = append(keys, k)
	}
	var keyMap = make(map[string][]string)
	keyMap["categories"] = keys
	return c.JSON(keyMap)
}

func getAllCards(c *fiber.Ctx) error {
	return c.JSON(cards[:9])
}

func getCard(c *fiber.Ctx) error {
	return c.SendString("single card")
}

func getAllManufacturers(c *fiber.Ctx) error {
	return c.SendString("all manufacturers")
}

func getManufacturer(c *fiber.Ctx) error {
	return c.SendString("single manufacturer")
}

func getAllPlayers(c *fiber.Ctx) error {
	return c.SendString("all players")
}

func getPlayer(c *fiber.Ctx) error {
	return c.SendString("single player")
}

func getAllTeams(c *fiber.Ctx) error {
	return c.SendString("all teams")
}

func getTeam(c *fiber.Ctx) error {
	return c.SendString("single team")
}
