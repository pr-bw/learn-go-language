package main

import (
	"errors"
	"fmt"
)

func pembagian(bilangan rune, pembagi rune) (rune, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi adalah 0")
	} else {
		return bilangan / pembagi, nil
	}
}

func main() {
	result, err := pembagian(20, 0)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(result)
	}
}
