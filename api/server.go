package api

import (
	"fmt"

	"github.com/VictorKabata/quotes-api/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var err error

func Run() {
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("Error fetching enviromental variable: %s\n", err)
	} else {
		fmt.Println("Successfully fetched enviromental variables")
	}

	server.InitializeServer()
	server.Run(":8081")

}
