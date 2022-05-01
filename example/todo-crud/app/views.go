package app

import (
	"gongo/core"
	"gongo/db"

	"strconv"
)

func HomePage(req *core.Request) core.Response {
	return core.HttpResponse("Welcome to Gongo!")
}

func AddUser(req *core.Request) core.Response {
	user := User{
		Name: "Shivendu", Email: "shivendu@foobar.com", Password: "XXXX##XXXX",
	}
	db.Conn.Create(&user)
	return core.HttpResponse("Created user " + user.Name)
}

func ListUsers(req *core.Request) core.Response {
	var users []User
	db.Conn.Find(&users)

	return core.JsonResponse(users)
}

func RemoveUser(req *core.Request) core.Response {
	var Id, err = strconv.Atoi(req.URL.Query()["Id"][0])
	if err != nil {
		return core.HttpResponse("Provide User Id")
	}
	db.Conn.Delete(&User{}, Id)

	return core.HttpResponse("Deleted successfully")
}

func UpdateUser(req *core.Request) core.Response {
	var Name = req.URL.Query()["Name"][0]
	var Id, err = strconv.Atoi(req.URL.Query()["Id"][0])
	if err != nil {
		return core.HttpResponse("Provide new user name")
	}
	db.Conn.Model(&User{}).Where("Id = ?", Id).Update("Name", Name)
	return core.HttpResponse("User name updated successfully")
}

func Hello(req *core.Request) core.Response {
	return core.HttpResponse("Hello world")
}

func Bye(req *core.Request) core.Response {
	return core.HttpResponse("Bye world")
}
