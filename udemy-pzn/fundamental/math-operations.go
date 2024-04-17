package main

import "fmt"

func main() {
	var result = (10 * 3 % 3) + 10/2

	fmt.Println(result)

	// augmented operator
	var i int8 = 10
	fmt.Println(i)

	i += 20
	fmt.Println(i)

	i -= 5
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

	i--
	fmt.Println(i)
}
