package main

import "fmt"

type Biodata struct {
	name, nationality string
	yob               uint16
}

func main() {
	personOne := Biodata{
		name:        "Aditya",
		nationality: "Indonesia",
		yob:         2002,
	}

	// passing by value
	personTwo := personOne
	personTwo.name = "Haruka"

	fmt.Println(personOne)
	fmt.Println(personTwo)

	fmt.Println("======")

	// passing by reference
	personThree := Biodata{
		name: "Ronaldo", nationality: "Portugal", yob: 1980,
	}

	fmt.Println(personThree)

	var personFour *Biodata = &personThree
	personFour.name = "Obama"

	fmt.Println(personThree)
	fmt.Println(personFour)

}
