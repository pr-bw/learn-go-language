package main

import "fmt"

type CustomError struct {
	ErrName string
	ErrCode int
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Error: %s, code: %d", c.ErrName, c.ErrCode)
}

func processInput(number int) error {
	if number < 0 {
		return &CustomError{"Number should be non-negative", 400}
	}

	return nil
}

func main() {
	err := processInput(3)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success")
	}
}
