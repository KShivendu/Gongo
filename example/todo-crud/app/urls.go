package app

import (
	"gongo/core"
)

var UrlPatterns = []core.Path{
	{Url: "/addUser", View: AddUser},       // C
	{Url: "/listUsers", View: ListUsers},   // R
	{Url: "/updateUser", View: UpdateUser}, // U
	{Url: "/removeUser", View: RemoveUser}, // R

	{Url: "/admin/login", View: core.AdminLogin}, // AdminLogin
	{Url: "/admin", View: core.AdminHome},        // AdminHome

	{Url: "/hello", View: Hello}, // Dummy
	{Url: "/bye", View: Bye},     // Dummy
	{Url: "/", View: HomePage},   // Landing Page
}
