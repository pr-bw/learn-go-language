package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Loop ke-", i)
	}

	fmt.Println()

	for j := 1; j <= 15; j++ {
		if j%2 == 0 {
			continue
		} else {
			fmt.Println(j)
		}
	}
}
