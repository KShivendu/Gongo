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
