package server

import (
	_ "log"

	"github.com/gofiber/fiber/v2"

	gm "gamch1k/spy-api/api/game_manager"
)


func Start(address string) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Spy game API")
	})


	app.Post("/api/new_game", func(c *fiber.Ctx) error {
		gm.NewGame()
		return c.SendString("Welcome to the Spy game API")
	})


	app.Get("/api/get_games", func(c *fiber.Ctx) error {
		return c.JSON(gm.Games)
	})


	app.Post("/api/connect/:game_id/:player_id", func(c *fiber.Ctx) error {
		game_id := c.Params("game_id")
		player_id := c.Params("player_id")
		return c.SendString("Game Id: " + game_id + "\nPlayer id: " + player_id)
	})

	app.Listen(address)
}
