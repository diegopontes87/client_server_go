package database

import (
	"context"
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
		log.Fatal("Error trying connect to the database:", err)
	}
	createTables()
}

func createTables() {
	creatingTableError := dbInstance.AutoMigrate(&entities.DBCotation{})
	if creatingTableError != nil {
		log.Fatal("Error while migration:", creatingTableError)
	}
}

func InsertCotation(ctx context.Context, cotation *entities.ServerCotation) {
	dbCotation := cotation.ConvertToDBCotation()

	select {
	case <-ctx.Done():
		log.Println("Timeout or cancellation occurred before starting the database operation")
		return
	default:
	}

	result := dbInstance.WithContext(ctx).Create(&dbCotation)
	if result.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout occurred while inserting data into the database")
		} else {
			log.Println("Error while inserting data:", result.Error)
		}
	} else {
		log.Println("Success inserting data")
	}
}
