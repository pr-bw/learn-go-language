package main

import "fmt"

func validatePassword(password string) map[string]string {
	if password == "" {
		return nil
	} else {
		return map[string]string{"password": password}
	}
}

func main() {
	var createPassword map[string]string = validatePassword("yuzhongmlbb10")

	if createPassword != nil {
		fmt.Println(createPassword["password"])
	} else {
		fmt.Println("Password is empty")
	}
}
