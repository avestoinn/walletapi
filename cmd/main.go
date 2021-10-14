package main

import (
	"fmt"
	"walletapi/pkg/database"
	"walletapi/pkg/server"
)

func main() {
	// Creating tables and making auto-migrate
	err := database.CreateTables()
	if err != nil {
		fmt.Printf("Can't create tables: %s", err)
	}

	// Running the server
	err = server.Run()
	if err != nil {
		fmt.Println(err)
	}
}
