package main

import (
	"fmt"
	"log"

	"cophee.team/rpts/config"
	"cophee.team/rpts/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn := database.ConnectToDatabase(config)
	defer conn.Close()

	database.ShowAllTables(conn)
	ress := database.HoursOfEmployeePerProject(conn)
	for _, id := range (ress) {
		fmt.Printf("%v\n", id)
	}
}