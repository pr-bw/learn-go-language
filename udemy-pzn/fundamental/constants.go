package main

import "fmt"

func main() {
	const fullName string = "Uzumaki Naruto"
	const unusedConstant = "Unused Constant"

	const (
		village = "Konohagakure"
		strong  = true
		spouse  = "Hinata"
	)

	fmt.Println(fullName)
	fmt.Println(village)
	fmt.Println(strong)
	fmt.Println(spouse)
}
