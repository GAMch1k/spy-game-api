# About

Go-To-Do-App is simple mobile and desktop app written on Go that work with [Go-To-Do-Api](https://github.com/GAMch1k/go-to-do-api), it can show, edit, delete and post tasks on server.

It was written just to learn more about Go.


# API endpoints

All endpoints should start with:

    http://192.168.178.20:3000


Create new game: 

    Request type POST
    /api/new_game

Get games:

    Request type GET
    /api/get_games

Connect to the game:

    Request type POST
    /api/connect/{game_id}/{player_id}