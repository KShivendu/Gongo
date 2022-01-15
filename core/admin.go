package core

type TodoItem struct {
	Name string
	Done bool
}

type TodoList struct {
	User string
	List []TodoItem
}

func AdminHome(req *Request) string {
	s := &TodoList{
		User: "KShivendu",
		List: []TodoItem{
			{Name: "Task 1", Done: false},
			{Name: "Task 2", Done: true},
		},
	}

	template := &Template{filepath: "index.html"}

	return template.render(s, req)
}

func AdminLogin(req *Request) string {
	template := &Template{filepath: "admin/login.html"}

	return template.render(nil, req)
}
