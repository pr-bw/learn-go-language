package main

import (
	"fmt"
	"fundamental/database"
	_ "fundamental/internal"
)

func main() {
	dbConnection := database.GetDatabase()
	fmt.Println(dbConnection)
}
