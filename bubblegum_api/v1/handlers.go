package bubblegum_api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

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
	skip := 0
	limit := 9
	var err error
	results := []Card{}
	if c.Query("skip") != "" {
		skip, err = strconv.Atoi(c.Query("skip"))
		if err != nil {
			return fmt.Errorf("could not convert id parameter to int: %v", err)
		}
	}
	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return fmt.Errorf("could not convert limit parameter to int: %v", err)
		}
	}

	if c.Query("search") != "" {
		text := c.Query("search")
		for _, card := range cards {
			if strings.Contains(strings.ToLower(card.Player), strings.ToLower(text)) ||
				strings.Contains(strings.ToLower(card.Description), strings.ToLower(text)) {
				results = append(results, card)
			}
		}
	}

	if skip > len(cards) || limit > len(cards) {
		return fiber.NewError(fiber.StatusNotFound, "Not Found - Query parameters exceed inventory")
	}

	// if searching, return this
	if len(results) > 0 {
		if skip < len(results) && skip+limit > len(results) {
			if len(results) > 0 {
				return c.JSON(results[skip:])
			}
			return c.JSON(results[skip:])
		}

		return c.JSON(results[skip : skip+limit])
	}

	// if NOT searching, default return
	if skip < len(cards) && skip+limit > len(cards) {
		if len(results) > 0 {
			return c.JSON(results[skip:])
		}
		return c.JSON(cards[skip:])
	}

	return c.JSON(cards[skip : skip+limit])
}

func getCard(c *fiber.Ctx) error {
	index, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("could not convert id parameter to int: %v", err)
	}
	if index > len(cards) {
		return fiber.NewError(fiber.StatusNotFound, "Card not found")
	}
	return c.JSON(cards[index-1])
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
