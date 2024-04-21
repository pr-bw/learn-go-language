package main

import "fmt"

func sumAll(args ...int) int {
	total := 0

	for _, arg := range args {
		total += arg
	}

	return total
}

func main() {
	// 1.) pass number args to parameter and considered as a slice
	total := sumAll(5, 5, 5, 5)
	fmt.Println(total)

	// 2.) pass a slice args to parameter
	var favNumbers = []int{7, 23, 32, 13, 5}

	// convert slice to varargs
	total = sumAll(favNumbers...)
	fmt.Println(total)
}
