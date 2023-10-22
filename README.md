# About

Spy-Game-API is an Golang API for Spy Game that I write with my friend, just for fun


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