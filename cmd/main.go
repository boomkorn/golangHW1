package main

import (
	"fmt"
	"golanghw1"
	"os"
)

func main() {
	us := golanghw1.UserService{"users.json"}

	if len(os.Args) <= 1 {
		fmt.Println("Please input with following function")
		fmt.Println("- add name={string} age={int}")
		fmt.Println("- list")
		fmt.Println("- clear")
		fmt.Println("- remove")
	} else {
		switch os.Args[1] {
		case "add":
			us.AddUser(os.Args[2], os.Args[3])
		case "list":
			us.PrintUser()
		case "clear":
			us.ClearUser()
		case "remove":
			us.RemoveUser()
		default:
			fmt.Println("Invalid method")
		}
	}
}
