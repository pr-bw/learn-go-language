package main

import "fmt"

func facRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * facRecursive(n-1)
	}
}

func main() {
	// 1.) factorial using for loop
	n := 1
	total := 1

	for i := 5; i > n; i-- {
		total *= i
	}

	fmt.Println(total)

	// 2.) factorial using recursive
	recResult := facRecursive(5)
	fmt.Println(recResult)
}
