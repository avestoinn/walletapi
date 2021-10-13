package main

import (
	"fmt"
	"walletapi/pkg/database"
)

func main() {
	// Creating tables and making auto-migrate
	err := database.CreateTables()
	if err != nil {
		fmt.Printf("Can't create tables: %s", err)
	}
}
