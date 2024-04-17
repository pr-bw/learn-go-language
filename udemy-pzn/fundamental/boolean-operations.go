package main

import "fmt"

func main() {
	var nilaiAkhir uint8 = 96
	var nilaiPresensi uint8 = 80

	// and operator
	var lulus bool = (nilaiAkhir > 80 && nilaiPresensi > 85)

	fmt.Println(lulus)

	// or operator
	lulus = (nilaiAkhir > 80 || nilaiPresensi > 85)

	fmt.Println(lulus)

	// not operator
	var isThisWorldIsBeautiful bool = !false

	fmt.Println(isThisWorldIsBeautiful)
}
