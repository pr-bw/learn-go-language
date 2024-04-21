package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Aditya"
	lastName = "Prabowo"

	return firstName, middleName, lastName
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
