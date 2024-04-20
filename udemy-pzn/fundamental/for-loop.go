package main

import "fmt"

func main() {
	// 1.) loop from an array

	var colors [3]string

	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	var counter = 0

	for counter < len(colors) {
		fmt.Println(colors[counter])
		counter++
	}

	fmt.Println()

	var members = [3]string{
		"Hana",
		"John",
		"Scarlett",
	}

	// init stt		condition	  post stt
	for i := 0; i < len(members); i++ {
		fmt.Println(members[i])
	}

	// use for range
	for _, member := range members {
		fmt.Println(member)
	}

	fmt.Println()

	// 2.) loop from a slice
	cartoons := []string{
		"Naruto Shippuden",
		"SpongeBob Squarepants",
		"One Piece",
		"Jujutsu Kaisen",
		"Dora The Explorer",
	}

	for index, cartoon := range cartoons {
		fmt.Println("index:", index, ", value :", cartoon)
	}

	fmt.Println()

	// 3.) loop from a map
	var myFirstGuitar map[string]string = map[string]string{
		"manufacturer": "Yamaha",
		"merk":         "Pacifica PAC012",
		"color":        "White",
	}

	for key, value := range myFirstGuitar {
		fmt.Println(key, ":", value)
	}
}
