package main

import "fmt"

type Women struct {
	Name string
}

func (women *Women) Married() {
	women.Name = "Mrs. " + women.Name
}

func main() {
	myOshi := Women{"Haruka"}

	myOshi.Married()

	fmt.Println(myOshi.Name)
}
