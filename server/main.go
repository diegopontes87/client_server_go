package main

import (
	"client_server/server/api"
	"client_server/server/database"
)

// RUN DATABASE LOCALLY - sqlite3 database.db

func main() {
	database.InitDB()
	api.CreateServer()
}
