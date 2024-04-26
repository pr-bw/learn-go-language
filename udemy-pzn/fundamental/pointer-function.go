package main

import "fmt"

type Address struct {
	city, country string
}

func changeAddressCity(address *Address) {
	address.city = "Jayapura"
}

func main() {
	address := Address{}

	changeAddressCity(&address)

	fmt.Println(address)
}
