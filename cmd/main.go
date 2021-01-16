package main

import (
	"fmt"
	"golanghw1"
	"os"
	"strconv"
	"strings"
)

func main() {
	fs := golanghw1.FileStore{"users.json"}
	ul := golanghw1.GetUser(fs.ReadFile())

	if len(os.Args) <= 1 {
		fmt.Println("Invalid argument")
	} else {
		switch os.Args[1] {
		case "add":
			addUser(ul, fs)
		case "list":
			ul.Print()
		case "clear":
			fs.ClearFile()
		case "remove":
			fs.RemoveFile()
		default:
			fmt.Println("Invalid method")
		}
	}
}
func addUser(ul golanghw1.UserList, fs golanghw1.FileStore) {
	if checkAddArgument(os.Args) {
		nameContext := strings.Split(os.Args[2], "=")
		ageContext := strings.Split(os.Args[3], "=")
		name := nameContext[len(nameContext)-1]
		age, _ := strconv.Atoi(ageContext[len(ageContext)-1])

		ul.AddUser(golanghw1.User{len(ul.Users) + 1, name, age}, fs)
	} else {
		fmt.Println("Invalid argument")
	}
}
func checkAddArgument(osArg []string) bool {
	if len(osArg) != 4 {
		return false
	} else if s := strings.Split(osArg[2], "="); len(s) == 2 && s[0] != "name" {
		return false
	} else if s := strings.Split(osArg[3], "="); len(s) == 2 && s[0] != "age" {
		return false
	}
	return true
}
