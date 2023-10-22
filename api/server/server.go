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

	app.Listen(address)
}
