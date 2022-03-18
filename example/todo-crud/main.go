package main

import (
	"gongo/core"
	"gongo/db"
	"strconv"

	"log"
	"net/http"

	"gongo/example/todo-crud/app"
)

func initServer() {
	for _, path := range app.UrlPatterns {
		_path := path // 'path' gets overriden and closure remembers only the pointer. So store the value before it gets changed in the next loop
		handleFunction := func(w http.ResponseWriter, req *http.Request) {
			var view = _path.View
			if req.URL.Path != _path.Url {
				view = core.HttpNotFoundHandler
			}
			response := view(
				&core.Request{Request: *req, Writer: &w, Args: core.KWA{}},
			)
			response.Write(w)
		}
		http.HandleFunc(_path.Url, handleFunction)
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
