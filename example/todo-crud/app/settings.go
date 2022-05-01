package app

import "gongo/db"

var Port int = 5000
var Debug bool = true

var AllowedHosts = []string{
	"",
}

var InstalledApps = []string{
	"core",
}

var Middlewares = []string{
	"AuthMiddleware",
}

var Databases = []db.Database{
	{
		Alias:  "default",
		Engine: "sqlite3",
		Name:   "BASE_DIR" + "db.sqlite3",
	},
}

var StaticUrl = "/static/"
var StaticDir = "static"
