package main

import "fmt"

type Bio struct {
	name, nationality string
	age               uint8
}

func main() {
	personOne := Bio{
		name:        "Prabowo",
		nationality: "Indonesia",
		age:         20,
	}

	fmt.Println(personOne)

	personTwo := &personOne
	personTwo.name = "Saitama"

	fmt.Println(personOne)
	fmt.Println(personTwo)

	//personTwo = &Bio{name: "Messi", nationality: "Argentina", age: 33}
	//
	//fmt.Println(personOne)
	//fmt.Println(personTwo)

	*personTwo = Bio{"Maeda", "Japan", 29}
	fmt.Println(personOne)
	fmt.Println(personTwo)
}
