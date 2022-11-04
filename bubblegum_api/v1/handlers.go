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

// GET the API resources
func getResources(c *fiber.Ctx) error {
	resources := make(map[string]string)
	resources["cards"] = "/api/v1/cards/"
	resources["manufacturers"] = "/api/v1/manufacturers/"
	resources["players"] = "/api/v1/players/"
	resources["teams"] = "/api/v1/teams/"
	resources["categories"] = "/api/v1/categories"
	return c.JSON(resources)
}

// GET all the baseball card categories
func getCategories(c *fiber.Ctx) error {
	keys := []string{}
	for k := range categories {
		keys = append(keys, k)
	}
	var keyMap = make(map[string][]string)
	keyMap["categories"] = keys
	return c.JSON(keyMap)
}

// GET all the baseball cards
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

	// return logic
	if c.Query("search") != "" && c.Query("category") == "" {
		// search cards
		text := c.Query("search")
		for _, card := range cards {
			if strings.Contains(strings.ToLower(card.Player), strings.ToLower(text)) ||
				strings.Contains(strings.ToLower(card.Description), strings.ToLower(text)) {
				results = append(results, card)
			}
		}
		if skip >= len(results) {
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		} else if skip < len(results) && skip+limit > len(results) {
			return c.JSON(results[skip:])
		} else {
			return c.JSON(results[skip : skip+limit])
		}
	} else if c.Query("search") != "" && c.Query("category") != "" {
		// search cards
		// filter by category
		text := c.Query("search")
		category := c.Query("category")
		for _, card := range cards {
			if strings.Contains(strings.ToLower(card.Player), strings.ToLower(text)) ||
				strings.Contains(strings.ToLower(card.Description), strings.ToLower(text)) {
				if card.Category == category {
					results = append(results, card)
				}
			}
		}
		if skip >= len(results) {
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		} else if skip < len(results) && skip+limit > len(results) {
			return c.JSON(results[skip:])
		} else {
			return c.JSON(results[skip : skip+limit])
		}
	} else if c.Query("category") != "" {
		// filter cards for category
		category := c.Query("category")
		for _, card := range cards {
			if card.Category == category {
				results = append(results, card)
			}
		}
		if skip >= len(results) {
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		} else if skip < len(results) && skip+limit > len(results) {
			return c.JSON(results[skip:])
		} else {
			return c.JSON(results[skip : skip+limit])
		}
	} else {
		// default return
		if skip >= len(cards) {
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		} else if skip < len(cards) && skip+limit > len(cards) {
			return c.JSON(cards[skip:])
		} else {
			return c.JSON(cards[skip : skip+limit])
		}
	}
}

// GET a single baseball card
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

// GET all card manufacturers
func getAllManufacturers(c *fiber.Ctx) error {
	skip := 0
	limit := 9
	var err error

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

	if skip < len(manufacturers) && skip+limit > len(manufacturers) {
		return c.JSON(manufacturers[skip:])
	} else if skip >= len(manufacturers) {
		return fiber.NewError(fiber.StatusNotFound, "Manufacturer not found")
	} else {
		return c.JSON(manufacturers[skip : skip+limit])
	}
}

// GET a single card manufacturer
func getManufacturer(c *fiber.Ctx) error {
	index, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("could not convert id parameter to int: %v", err)
	}
	if index > len(manufacturers) {
		return fiber.NewError(fiber.StatusNotFound, "Manufacturer not found")
	}
	return c.JSON(manufacturers[index-1])
}

// GET all players
func getAllPlayers(c *fiber.Ctx) error {
	skip := 0
	limit := 9
	var err error

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

	if skip < len(players) && skip+limit > len(players) {
		return c.JSON(players[skip:])
	} else if skip >= len(players) {
		return fiber.NewError(fiber.StatusNotFound, "Cards not found")
	} else {
		return c.JSON(players[skip : skip+limit])
	}
}

// GET single player
func getPlayer(c *fiber.Ctx) error {
	index, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("could not convert id parameter to int: %v", err)
	}
	if index > len(players) {
		return fiber.NewError(fiber.StatusNotFound, "Player not found")
	}
	return c.JSON(players[index-1])
}

// GET all teams
func getAllTeams(c *fiber.Ctx) error {
	skip := 0
	limit := 9
	var err error

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

	if skip < len(teams) && skip+limit > len(teams) {
		return c.JSON(teams[skip:])
	} else if skip >= len(teams) {
		return fiber.NewError(fiber.StatusNotFound, "Teams not found")
	} else {
		return c.JSON(teams[skip : skip+limit])
	}
}

// GET a single team by ID
func getTeam(c *fiber.Ctx) error {
	index, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("could not convert id parameter to int: %v", err)
	}
	if index > len(teams) {
		return fiber.NewError(fiber.StatusNotFound, "Team not found")
	}
	return c.JSON(teams[index-1])
}
