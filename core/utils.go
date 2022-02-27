package core

import (
	"encoding/json"
	"io"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Response struct {
	Headers []Header
	Body    string
}

type Header struct {
	Name  string
	Value string
}

func (res Response) Write(w http.ResponseWriter) {
	for _, header := range res.Headers {
		w.Header().Set(header.Name, header.Value)
	}
	io.WriteString(w, res.Body)
}

func HttpResponse(body string) Response {
	headers := []Header{{"Content-Type", "text/html"}}
	e := Response{Headers: headers, Body: body}
	return e
}

func JsonResponse(body interface{}) Response {
	headers := []Header{{"Content-Type", "application/json"}}
	response, _ := json.Marshal(body)
	e := Response{Headers: headers, Body: string(response)}
	return e
}
