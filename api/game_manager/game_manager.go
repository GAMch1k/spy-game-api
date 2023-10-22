package game_manager

import (
	_ "math/rand"
	_ "log"

	"gamch1k/spy-api/api/database"
)

type Game struct {
	GameId    int       `json:"game_id"` 
	Players   []string  `json:"players"`
	SpyIds    []int     `json:"spy_ids"`
	Topic     string    `json:"topic"`
	Started   bool      `json:"started"`
}

var Games = []Game{}


func GetGames() *[]Game {
	return &Games
}


func GameExists(gameId int) bool {
	for _, el := range Games {
		if el.GameId == gameId {
			return true
		}
	}
	
	return false
}


func NewGame() {
	database.CreateGame()
}

func Connect(game_id int, player_id string) (bool, error) {
	
	if !GameExists(game_id) {
		return false, nil
	}

	for _, game := range *GetGames() {
		if game.GameId == game_id {
			
			game.Players = append(game.Players, player_id)
		}
	}
	
	
	return true, nil
}

