package main

import "fmt"

func myAnyFunc() interface{} {
	// return 2
	// return true
	// return 2.5
	return !false
}

func main() {
	var myVar any = myAnyFunc()

	fmt.Println(myVar)
}
