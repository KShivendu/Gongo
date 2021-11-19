package main

import (
	"gongo/core"
	"gongo/db"
	"io"
	"strconv"

	"log"
	"net/http"
)

func initServer() {
	for _, path := range core.UrlPatterns {
		_path := path // 'path' gets overriden and closure remembers only the pointer. So store the value before it gets changed in the next loop
		http.HandleFunc(_path.Url, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, _path.View(
				&core.Request{Req: req, Args: core.KWA{}},
			))
		})
	}

	var port = ":" + strconv.Itoa(core.Port)
	log.Println("Listing for requests at http://localhost" + port + "/hello")
	log.Fatal(http.ListenAndServe(port, nil))
}

func initDB() {
	db.Setup()
	db.Conn.AutoMigrate(&core.User{})
}

func main() {
	initDB()
	initServer()
}
