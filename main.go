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
		_path := path // 'path' gets overriden and closure remembers only the pointer. So store the value before it gets changed in the next loop
		http.HandleFunc(_path.Url, func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, _path.View(req))
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
