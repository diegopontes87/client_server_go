package main

import (
	"client_server/server/api"
	"client_server/server/database"
)

func main() {
	database.InitDB()
	api.CreateServer()
}
