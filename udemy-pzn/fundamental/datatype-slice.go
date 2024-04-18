package main

import "fmt"

func main() {
	// 1.) create slice from an array

	var favAnimeCharacter = [5]string{
		"Shikamaru", "Zoro", "Gojo", "Nobita", "Goku",
	}

	var sliceOne = favAnimeCharacter[2:5]

	fmt.Println(favAnimeCharacter)
	fmt.Println(sliceOne, "\n")

	sliceOne[1] = "Newbieta"

	fmt.Println(favAnimeCharacter)
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne), "\n")

	sliceTwo := append(sliceOne, "Moskov") // creates a new array because the previous slice capacity is full.
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)
	fmt.Println(favAnimeCharacter, "\n")

	// 2.) create slice directly
	myFavGuitars := []string{"Fender", "Gibson", "ESP", "PRS", "Yamaha"}

	var herFavGuitars []string = make([]string, len(myFavGuitars), cap(myFavGuitars))

	// just copy the value and not the reference
	copy(herFavGuitars, myFavGuitars)

	fmt.Println(myFavGuitars)
	fmt.Println(herFavGuitars)

	myFavGuitars[1] = "Squire"

	fmt.Println(myFavGuitars)
	fmt.Println(herFavGuitars)
}
