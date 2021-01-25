package service

import "fmt"

func Debug(text string, obj interface{}) {
	fmt.Println("*********************************")
	fmt.Printf("%s : %v\n", text, obj)
	fmt.Println("/////////////////////////////////")
}
