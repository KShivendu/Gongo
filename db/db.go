package db

import (
	"fmt"
	"log"
	"strconv"

	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Alias    string
	Engine   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

// Db is the database connection
var Conn *gorm.DB

func Setup(databaseConfig Database) {
	if databaseConfig.Engine == "sqlite3" {
		var err error
		Conn, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else if databaseConfig.Engine == "postgres" {
		var err error
		dsn := fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			databaseConfig.Host, strconv.Itoa(databaseConfig.Port),
			databaseConfig.Name,
			databaseConfig.Username, databaseConfig.Password,
		)
		Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	log.Println("Connected to the DB")
}
