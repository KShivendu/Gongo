package core

func HttpNotFoundHandler(req *Request) Response {
	template := &Template{filepath: "404.html"}
	var context = map[string]interface{}{
		"core.RequestURL": "http://127.0.0.1:8000/",
		"PossibleUrlPatterns": []interface{}{
			"Boom/", "admin/",
		},
	}

	return template.render(context)
}
