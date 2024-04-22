package main

import "fmt"

func main() {
	myInitNumber := 1

	increment := func() {
		fmt.Println("increment")
		myInitNumber++
	}

	increment()
	increment()
	increment()

	fmt.Println(myInitNumber)
}
