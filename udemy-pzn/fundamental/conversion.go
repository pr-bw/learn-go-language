package main

import "fmt"

func main() {
	// number data types conversion
	var A int32 = 4000
	var B int64 = int64(A)
	var C int8 = int8(A)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	// integer to string data types conversion
	email := "freyajkt48@gmail.com"
	var eByte uint8 = email[2]
	var eString = string(eByte)

	fmt.Println(eByte)
	fmt.Println(eString)
}
