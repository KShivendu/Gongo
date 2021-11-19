package db

import (
	"log"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Db is the database connection
var Conn *gorm.DB

func Setup() {
	var err error
	Conn, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to the DB")
}
