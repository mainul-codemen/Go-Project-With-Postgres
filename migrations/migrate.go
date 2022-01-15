package main

import (
	"Go-Project-With-Postgres/storage/postgres"
	"log"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
