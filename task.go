package golanghw1

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type UserList struct {
	Users []User
}

// User management
func GetUser(byteValue []byte) UserList {
	var userList UserList
	json.Unmarshal(byteValue, &userList.Users)
	return userList
}
func (ul UserList) AddUser(u User, fs FileStore) {
	ul.Users = append(ul.Users, u)
	fs.WriteFile(ul)
}
func (ul *UserList) PrintUser() {
	if len(ul.Users) == 0 {
		fmt.Printf("{}")
	} else {
		jsonString, _ := json.MarshalIndent(ul.Users, "", "    ")
		fmt.Printf("%s", jsonString)
	}
}
