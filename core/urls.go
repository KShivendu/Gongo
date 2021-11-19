package core

import (
	"net/http"
)

type KWA map[string]interface{}
type Request struct {
	Req  *http.Request
	Args KWA
}

type View func(req *Request) string

type Path struct {
	Url  string
	View View
}

var UrlPatterns = []Path{
	{Url: "/addUser", View: AddUser},       // C
	{Url: "/listUsers", View: ListUsers},   // R
	{Url: "/updateUser", View: UpdateUser}, // U
	{Url: "/removeUser", View: RemoveUser}, // R
	{Url: "/hello", View: Hello},
	{Url: "/bye", View: Bye},
}
