package main

import (
	"gongo/core"
	"gongo/db"
	"gongo/example/todo-crud/app"
	"strings"

	"log"
	"net/http"
	"strconv"
)

func initServer() {
	http.Handle(app.StaticUrl, http.StripPrefix(app.StaticUrl, http.FileServer(http.Dir(app.StaticDir))))

	for _, path := range app.UrlPatterns {
		_path := path // 'path' gets overriden and closure remembers only the pointer. So store the value before it gets changed in the next loop
		handleFunction := func(w http.ResponseWriter, req *http.Request) {
			var view = _path.View
			if !strings.HasPrefix(req.URL.Path, app.StaticUrl) && req.URL.Path != _path.Url {
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
	log.Println("Listing for requests at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initDB() {
	// db.Setup(db.Database{Alias: "default", Engine: "sqlite3", Name: "BASE_DIR" + "cool.sqlite3"})
	db.Setup(db.Database{Alias: "psql", Engine: "postgres", Host: "localhost", Port: 5432, Name: "gongo", Username: "postgres", Password: "something"})
	db.Conn.AutoMigrate(&core.User{})
}

func main() {
	initDB()
	initServer()
}
