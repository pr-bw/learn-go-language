package main

import "fmt"

func main() {
	var myFavLanguages [3]string

	myFavLanguages[0] = "Japanese"
	myFavLanguages[1] = "English"
	myFavLanguages[2] = "Indonesian"

	fmt.Println(myFavLanguages)

	// array function
	myFavLanguages[2] = "Jawa"

	fmt.Println(myFavLanguages)

	var myFavNumbers = [3]uint8{
		23,
		32,
		7,
	}

	fmt.Println(myFavNumbers)

	var myCats = [...]string{
		"Titi",
		"Mei-Mei",
		"Jeje",
		"Moi",
		"Mona",
		"Momo",
	}

	// array function
	fmt.Println(len(myCats))
	fmt.Println(myCats[2])
}
