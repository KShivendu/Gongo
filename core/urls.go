package core

import (
	"net/http"
)

type KWA map[string]interface{}
type Request struct {
	http.Request
	Writer *http.ResponseWriter
	Args   KWA
}

type View func(req *Request) Response

type Path struct {
	Url  string
	View View
}

var UrlPatterns = []Path{
	{Url: "/addUser", View: AddUser},       // C
	{Url: "/listUsers", View: ListUsers},   // R
	{Url: "/updateUser", View: UpdateUser}, // U
	{Url: "/removeUser", View: RemoveUser}, // R

	{Url: "/admin/login", View: AdminLogin}, // AdminLogin
	{Url: "/admin", View: AdminHome},        // AdminHome

	{Url: "/hello", View: Hello}, // Dummy
	{Url: "/bye", View: Bye},     // Dummy
	{Url: "/", View: HomePage},   // Landing Page
}
