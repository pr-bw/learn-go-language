package main

import "fmt"

func sayGoodbyeTo(firstName string, lastName string) {
	fmt.Printf("Goodbye %s %s", firstName, lastName)
}

func main() {
	sayGoodbyeTo("Hatsune", "Miku")
}
