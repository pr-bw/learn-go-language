package main

import "fmt"

func main() {
	var foodName string = "Bakso"

	switch foodName {
	case "Sate":
		fmt.Println("Sate itu enak bro")
	case "Bakso":
		fmt.Println("Bakso itu enak bro")
	case "Hamburger":
		fmt.Println("Hamburger itu enak bro")
	default:
		fmt.Println("Unknown food name")
	}

	// short switch

	switch length := len(foodName); length < 7 {
	case true:
		fmt.Println("Jumlah karakter dari food name lebih sedikit dari 7")
	case false:
		fmt.Println("Jumlah karakter dari food name lebi banyak dari 7")
	}
}
