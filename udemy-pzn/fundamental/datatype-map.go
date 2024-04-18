package main

import "fmt"

func main() {
	// the first way to create map
	var myMotor map[string]string = make(map[string]string)

	myMotor["merk"] = "Satria FU Predator 2017"
	myMotor["color"] = "Black"
	myMotor["manufacturer"] = "Suzuki"
	myMotor["fuel"] = "Oil"

	// get data by map key
	fmt.Println(myMotor["color"])

	// get all pair key-value data from map
	fmt.Println(myMotor)

	// delete pair key-value data from map by key
	delete(myMotor, "fuel")

	fmt.Println(myMotor)

	// second way to make map

	myCuteCat := map[string]string{
		"name":       "Mona",
		"skin_color": "White",
		"gender":     "Male",
		"mom":        "Titi",
	}

	fmt.Println(myCuteCat)
}
