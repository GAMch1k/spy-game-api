package main

import (
	"gamch1k/spy-api/api/server"
	"gamch1k/spy-api/api/database"
)


func main() {
	database.InitDatabase()
	server.Start("0.0.0.0:3000")
}
