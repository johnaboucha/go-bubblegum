package bubblegum_api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/", apiHome)
	app.Get("/api/v1/resources/", getResources)
	app.Get("/api/v1/categories/", getCategories)

	app.Get("/api/v1/cards/", getAllCards)
	app.Get("/api/v1/cards/:id", getCard)
	// app.Post("/api/v1/cards/", addCards)
	// app.Put("/api/vi/cards/:id", updateCard)
	// app.Delete("/api/vi/cards/:id", deleteCard)

	app.Get("/api/v1/manufacturers/", getAllManufacturers)
	app.Get("/api/v1/manufacturers/:id", getManufacturer)

	app.Get("/api/v1/players/", getAllPlayers)
	app.Get("/api/v1/players/:id", getPlayer)

	app.Get("/api/v1/teams/", getAllTeams)
	app.Get("/api/v1/teams/:id", getTeam)
}
