package main

import "fmt"

type Greet func(string) string

func greetHer(name string, greet Greet) string {
	return greet(name)
}

func main() {
	// 1.) example one
	square := func(x uint8) uint8 {
		return x * x
	}

	squareResult := square(10)

	fmt.Println(squareResult)

	// 2.) example two
	greet := greetHer("Halmahera", func(name string) string {
		if len(name) > 5 {
			return "Hi, " + name
		} else {
			return "Hello, " + name
		}
	})

	fmt.Println(greet)
}
