package core

import (
	"encoding/json"
	"gongo/db"
	"strconv"
)

func AddUser(req *Request) string {
	user := User{
		Name: "Shivendu", Email: "shivendu@foobar.com", Password: "XXXX##XXXX",
	}
	db.Conn.Create(&user)
	return "Created user " + user.Name
}

func ListUsers(req *Request) string {
	var users []User
	db.Conn.Find(&users)

	response, _ := json.Marshal(users)
	return string(response)
}

func RemoveUser(req *Request) string {
	var Id, err = strconv.Atoi(req.Req.URL.Query()["Id"][0])
	if err != nil {
		return "Provide User Id"
	}
	db.Conn.Delete(&User{}, Id)
	return "Deleted successfully"
}

func UpdateUser(req *Request) string {
	var Name = req.Req.URL.Query()["Name"][0]
	var Id, err = strconv.Atoi(req.Req.URL.Query()["Id"][0])
	if err != nil {
		return "Provide new user name"
	}
	db.Conn.Model(&User{}).Where("Id = ?", Id).Update("Name", Name)
	return "User name updated successfully"
}

func Hello(req *Request) string {
	return "Hello world"
}

func Bye(req *Request) string {
	return "Bye world"
}
