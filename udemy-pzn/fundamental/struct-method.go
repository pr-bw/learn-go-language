package main

import "fmt"

type Athlete struct {
	Name, Nationality string
	Age               uint8
}

func (athlete Athlete) introduction(name string) {
	fmt.Printf("Hello %s, my name is %s.\n", name, athlete.Name)
}

func main() {
	firstAthlete := Athlete{
		Name:        "Usain Bolt",
		Nationality: "Jamaica",
		Age:         28,
	}

	firstAthlete.introduction("Prabowo")
}
