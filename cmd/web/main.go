package main

import (
	"tools/db"
	"tools/internal/route"
)

func main() {
	configs := db.Configs{
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     "3306",
		Database: "test",
	}
	db.Open(configs)
	route.Routes()
}
