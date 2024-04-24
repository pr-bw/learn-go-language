package main

import "fmt"

func myRandomFunction() any {
	//return "good luck"
	//return 12
	//return false
	return 3.14
}

func main() {
	myVariable := myRandomFunction()

	switch myVariable.(type) {
	case string:
		fmt.Println("string", myVariable)
	case int:
		fmt.Println("int", myVariable)
	case bool:
		fmt.Println("bool", myVariable)
	default:
		fmt.Println("Unknown type", myVariable)
	}
}
