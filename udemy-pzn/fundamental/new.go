package main

import "fmt"

type Customer struct {
	name, address string
}

func main() {
	cusOne := new(Customer)
	cusTwo := cusOne

	cusTwo.name = "Prabowo"
	cusTwo.address = "Jakarta"

	fmt.Println(cusOne)
	fmt.Println(cusTwo)
}
