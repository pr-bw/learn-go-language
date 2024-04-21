package main

import "fmt"

func barkDoggy(name string) string {
	return name + " guk guk!"
}

func main() {
	firstDog := barkDoggy
	fmt.Println(firstDog("Carmell"))

	secondDog := barkDoggy
	fmt.Println(secondDog("Luccie"))
}
