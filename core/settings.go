package core

var Port int = 5000
var Debug bool = true

var AllowedHosts = []string{
	"example.com",
}

var InstalledApps = []string{
	"core",
}

var Middlewares = []string{
	"AuthMiddleware",
}

type Database struct {
	Alias  string
	Engine string
	Name   string
}

var Databases = []Database{
	{
		Alias:  "default",
		Engine: "sqlite3",
		Name:   "BASE_DIR" + "db.sqlite3",
	},
}

var StaticUrl = "/static/"
