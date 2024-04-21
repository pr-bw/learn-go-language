package main

import "fmt"

func sayonara(name string) string {
	sayonara := "Sayonara, " + name + " !!"
	return sayonara
}

func main() {
	var result string = sayonara("Robin")
	fmt.Println(result)
}
