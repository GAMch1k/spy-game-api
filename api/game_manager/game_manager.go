package game_manager

import (
	"math/rand"
	"log"
)

type Game struct {
	GameId    int       `json:"game_id"` 
	Players   []string  `json:"players"`
	SpyIds    []int     `json:"spy_ids"`
	Topic     string    `json:"topic"`
	Started   bool      `json:"started"`
}

var Games = []Game{}


func GameExists(gameId int) bool {
	for _, el := range Games {
		if el.GameId == gameId {
			return true
		}
	}
	
	return false
}


func NewGame() {
	game_id := rand.Intn(99999)

	for GameExists(game_id) {
		game_id = rand.Intn(99999)
	}

	Games = append(Games, Game{
		GameId: game_id,
		Players: []string{},
		SpyIds: []int{},
		Topic: "",
		Started: false,
	})

	log.Println(Games)
}
