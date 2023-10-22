package game_manager

import (
	"math/rand"
	"log"
)

type Game struct {
	GameId    int      `json:"game_id"` 
	Players   []int    `json:"players"`
	SpyIds    []int    `json:"spy_ids"`
	Topic     string   `json:"topic"`
}

var Games = []Game{}





func NewGame() {
	Games = append(Games, Game{
		GameId: rand.Intn(99999),
		Players: []int{},
		SpyIds: []int{},
		Topic: "",
	})

	log.Println(Games)
}
