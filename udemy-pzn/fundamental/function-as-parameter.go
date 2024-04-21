package main

import "fmt"

type Modulo func(int, int) int

func applyModulo(x int, y int, modulo Modulo) int {
	return modulo(x, y)
}

func modulo(x int, y int) int {
	return x % y
}

type Filter func(string) string

func sayOneWord(word string, filter Filter) string {
	return filter(word)
}

func filterWord(word string) string {
	if word == "bangsatkau" {
		return "Mulut anda kotor, tidak pantas untuk menjadi seorang pemimpin!"
	} else {
		return word
	}
}

func main() {
	moduloResult := applyModulo(5, 2, modulo)
	fmt.Println(moduloResult)

	filterResult := sayOneWord("bangsatkau", filterWord)
	fmt.Println(filterResult)
}
