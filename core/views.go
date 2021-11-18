package core

import "net/http"

func Hello2(req *http.Request) string {
	return "Hello world"
}

func Hello(req *http.Request) string {
	return "Hello world"
}

func Bye(req *http.Request) string {
	return "Bye world"
}
