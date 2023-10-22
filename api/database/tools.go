package database

import (
	_ "log"
	_ "errors"

	_ "github.com/mattn/go-sqlite3"
)

type Game struct {
	GameId 		int 	`json:"game_id"`
	PlayerId	string 	`json:"player_id"`
	SpyId 		string 	`json:"spy_id"`
	Topic 		string 	`json:"topic"`
	Started 	int 	`json:"started"`
}


func CreateGame() (bool, error){
	db := OpenDatabase()
	defer CloseDatabase(db)

	insert_text := `INSERT INTO games DEFAULT VALUES`

	query, err := db.Prepare(insert_text)
	if err != nil {
		return false, err
	}

	_, err = query.Exec()
	if err != nil {
		return false, err
	}

	return true, nil
}


func GetGames() ([]Game, error) {
	db := OpenDatabase()
	defer CloseDatabase(db)

	row, err := db.Query("SELECT * FROM games")
	if err != nil {
		return []Game{}, err
	}

	defer row.Close()

	var final []Game

	for row.Next() {
		var game_id int 
		var player_id string 
		var spy_id string 
		var topic string 
		var started int 
		
		row.Scan(&game_id, &player_id, &spy_id, &topic, &started)

		final = append(final, Game{
			GameId: game_id,
			PlayerId: player_id,
			SpyId: spy_id,
			Topic: topic,
			Started: started,
		})
	}

	return final, nil 
}


func CheckIfGameExists(game_id int) (bool, error) {
	res, err := GetGameById(game_id)

	return res.GameId != 0, err
}


func GetGameById(game_id int) (Game, error) {
	db := OpenDatabase()
	defer CloseDatabase(db)

	var game Game
	rows, err := db.Query("SELECT * FROM games WHERE game_id = ?")

	if err != nil {
		return Game{}, nil
	}
	
	for rows.Next() {
		rows.Scan(&game.GameId, &game.PlayerId, &game.SpyId, &game.Topic, &game.Started)
	}

	return game, nil
}


