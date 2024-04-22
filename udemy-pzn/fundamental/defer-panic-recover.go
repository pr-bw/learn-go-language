package main

import "fmt"

func endApp() {
	fmt.Println("Program finished")

	errMessage := recover()

	fmt.Println("panic occurs with the message", errMessage)
}

func runApp(error bool) string {
	if error {
		panic("ERROR BRO")
	} else {
		return "Program started"
	}
}

func main() {
	defer endApp()

	statusApp := runApp(true)
	fmt.Println(statusApp)
}
