package main

import "fmt"

func main() {
	email := "vp.freya@gmail.com"

	if email == "mng.aditya@gmail.com" {
		fmt.Println("Hi, Manager Aditya")
	} else if email == "vp.freya@gmail.com" {
		fmt.Println("Hi, Vice President Freya")
	} else if email == "eng.hera@gmail.com" {
		fmt.Println("Hi, Engineer Hera")
	} else {
		fmt.Println("Your email is not registered, try to contact admin")
	}

	// short if statement

	if charLength := len(email); charLength >= 15 {
		fmt.Println("quite a lot")
	} else {
		fmt.Println("quite a bit")
	}
}
