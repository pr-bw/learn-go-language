package main

import "fmt"

func getCustomerName() (string, string) {
	return "Aditya", "Prabowo"
}

func main() {
	firstName, lastName := getCustomerName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	_, b := getCustomerName()
	fmt.Println(b)
}
