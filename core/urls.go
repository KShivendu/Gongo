package core

import (
	"net/http"
)

type View func(req *http.Request) string

type Path struct {
	Url  string
	View View
}

var UrlPatterns = []Path{
	{Url: "/hello", View: Hello},
	{Url: "/bye", View: Bye},
}
