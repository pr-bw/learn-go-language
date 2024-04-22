package main

import "fmt"

func main() {
	/*
		Komentar yang baik adalah pada kode itu sendiri,
		kode dibuat supaya readable
	*/

	/*
		this is multiline
		comments
	*/

	// this is single line comment

	// i love golang <3

	/*
		runDivision is variable function and have two parameters,
		the function can only receive two arguments,
		it has return keyword to return the value of result division between parameter x and y
	*/
	runDivision := func(x, y float32) float32 {
		return x / y
	}

	result := runDivision(10, 3)
	fmt.Println(result)
}
