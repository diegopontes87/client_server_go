package database

import (
	"log"

	"client_server/server/entities"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
)

func InitDB() {
	var err error

	dbInstance, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error trying to connect to the database:", err)
	}
	createTables()
}

func createTables() {
	resetingError := dbInstance.Unscoped().Delete(&entities.DBCotation{}).Error
	if resetingError != nil {
		log.Fatal("Error while resiting table:", resetingError)
	}

	creatingTableError := dbInstance.AutoMigrate(&entities.DBCotation{})
	if creatingTableError != nil {
		log.Fatal("Error while migration:", creatingTableError)
	}
}

func InsertCotation(cotation *entities.ServerCotation) {
	tx := dbInstance.Begin()
	if tx.Error != nil {
		log.Println("Error starting transaction:", tx.Error)
		return
	}
	dbCotation := cotation.ConvertToDBCotation()
	result := tx.Create(&dbCotation)
	if result.Error != nil {
		log.Println("Error while inserting data:", result.Error)
	} else {
		log.Println("Success inserting data")
		tx.Commit()
	}
}
