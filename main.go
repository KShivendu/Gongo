package main

import (
	core "gongo/core"
	"io"
	"strconv"

	"log"
	"net/http"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&core.User{})
}

func initServer() {
	for _, path := range core.UrlPatterns {
		http.HandleFunc(path.Url, func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, path.View(req))
		})
	}

	var port = ":" + strconv.Itoa(core.Port)
	log.Println("Listing for requests at http://localhost" + port + "/hello")
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	initDB()
	initServer()
}
