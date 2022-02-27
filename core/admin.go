package core

type TodoItem struct {
	Name string
	Done bool
}

type TodoList struct {
	User string
	List []TodoItem
}

func AdminHome(req *Request) Response {
	s := &TodoList{
		User: "KShivendu",
		List: []TodoItem{
			{Name: "Task 1", Done: false},
			{Name: "Task 2", Done: true},
		},
	}

	template := &Template{filepath: "index.html"}

	return template.render(s)
}

func AdminLogin(req *Request) Response {
	template := &Template{filepath: "admin/login.html"}

	return template.render(nil)
}
