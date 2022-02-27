package core

import (
	"gongo/db"
	"net/http"
	"strconv"
)

func HomePage(req *Request) Response {
	if req.URL.Path != "/" {
		http.NotFound(*req.Writer, &req.Request)
		return HttpResponse("")
	}

	return HttpResponse("Welcome to Gongo!")
}

func AddUser(req *Request) Response {
	user := User{
		Name: "Shivendu", Email: "shivendu@foobar.com", Password: "XXXX##XXXX",
	}
	db.Conn.Create(&user)
	return HttpResponse("Created user " + user.Name)
}

func ListUsers(req *Request) Response {
	var users []User
	db.Conn.Find(&users)

	return JsonResponse(users)
}

func RemoveUser(req *Request) Response {
	var Id, err = strconv.Atoi(req.URL.Query()["Id"][0])
	if err != nil {
		return HttpResponse("Provide User Id")
	}
	db.Conn.Delete(&User{}, Id)

	return HttpResponse("Deleted successfully")
}

func UpdateUser(req *Request) Response {
	var Name = req.URL.Query()["Name"][0]
	var Id, err = strconv.Atoi(req.URL.Query()["Id"][0])
	if err != nil {
		return HttpResponse("Provide new user name")
	}
	db.Conn.Model(&User{}).Where("Id = ?", Id).Update("Name", Name)
	return HttpResponse("User name updated successfully")
}

func Hello(req *Request) Response {
	return HttpResponse("Hello world")
}

func Bye(req *Request) Response {
	return HttpResponse("Bye world")
}
