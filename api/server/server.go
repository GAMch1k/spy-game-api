package server

import (
	_ "log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	gm "gamch1k/spy-api/api/game_manager"
	"gamch1k/spy-api/api/database"
)


func Start(address string) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Spy game API")
	})	


	app.Post("/api/new_game", func(c *fiber.Ctx) error {
		res, err := database.CreateGame()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if !res {
			return c.Status(fiber.StatusBadRequest).SendString("Something went wrong")
		}
		return c.Status(fiber.StatusCreated).SendString("New game created successfully!")
	})


	app.Get("/api/get_games", func(c *fiber.Ctx) error {
		res, err := database.GetGames()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.JSON(res)
	})


	app.Post("/api/connect/:game_id/:player_id", func(c *fiber.Ctx) error {
		game_id, err := strconv.Atoi(c.Params("game_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Game id is not an integer value!")
		}
		player_id := c.Params("player_id")

		res, err := gm.Connect(game_id, player_id)

		if err != nil {
			return c.Status(fiber.StatusBadGateway).SendString("Something went wrong!")
		}

		if !res {
			return c.Status(fiber.StatusBadRequest).SendString("Game does not exists!")
		}

		return c.Status(fiber.StatusAccepted).SendString("Game Id: " + c.Params("game_id") + "\nPlayer id: " + player_id)
	})

	app.Listen(address)
}
