package main

import (
	"fmt"

	"github.com/gustavofalcione/first-go-api/database"
	"github.com/gustavofalcione/first-go-api/routes"
)

func main() {
	database.ConectDatabase()
	fmt.Println("Server is running 🔥")
	routes.HandleRequest()
}
