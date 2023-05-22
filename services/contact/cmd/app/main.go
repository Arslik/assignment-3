package main

import (
	"assignment3/pkg/store/postgres"
	"fmt"
)

func main() {
	db, err := postgres.ConnectDb("localhost", "5432", "postgres", "1", "go_microservices")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}
	defer db.Close()
}
