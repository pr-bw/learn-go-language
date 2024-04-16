package main

import "fmt"

func main() {
	type NIK uint32

	var nik NIK = 33333338

	var contohNIK = 1111111

	var convert NIK = NIK(contohNIK)

	fmt.Println(nik)
	fmt.Println(convert)

}
