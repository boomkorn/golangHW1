package golanghw1

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type UserList struct {
	Users []User
}
type UserService struct {
	FileName string
}

// User management
func (us UserService) getUser() UserList {
	fs := FileStore{us.FileName}

	var userList UserList
	json.Unmarshal(fs.ReadFile(), &userList.Users)
	return userList
}
func (us UserService) AddUser(nameContext string, ageContext string) {
	ul := us.getUser()
	name, age, error := formattingAdd(nameContext, ageContext)
	if error != nil {
		fmt.Printf("%v", error.Error())
	} else {
		ul.Users = append(ul.Users, User{len(ul.Users) + 1, name, age})
		fs := FileStore{us.FileName}
		jsonByte, _ := json.Marshal(ul.Users)
		fs.WriteFile(jsonByte)
	}
}
func (us UserService) PrintUser() {
	ul := us.getUser()
	if len(ul.Users) == 0 {
		fmt.Printf("{}")
	} else {
		jsonString, _ := json.MarshalIndent(ul.Users, "", "    ")
		fmt.Printf("%s", jsonString)
	}
}
func (us UserService) ClearUser() {
	fs := FileStore{us.FileName}
	fs.ClearFile()
}
func (us UserService) RemoveUser() {
	fs := FileStore{us.FileName}
	fs.RemoveFile()
}

// Argument management
func formattingAdd(nameContext string, ageContext string) (string, int, error) {
	if s := strings.Split(nameContext, "="); len(s) == 2 && s[0] != "name" {
		return "", 0, errors.New("Parameter invalid")
	} else if s := strings.Split(ageContext, "="); len(s) == 2 && s[0] != "age" {
		return "", 0, errors.New("Parameter invalid")
	} else {
		nameContext := strings.Split(os.Args[2], "=")
		ageContext := strings.Split(os.Args[3], "=")
		name := nameContext[len(nameContext)-1]
		age, err := strconv.Atoi(ageContext[len(ageContext)-1])
		if err != nil {
			return "", 0, errors.New("Age type is only int, please try again")
		} else {
			return name, age, nil
		}
	}

}
