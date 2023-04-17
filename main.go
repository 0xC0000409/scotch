package main

import (
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	server.Init()
}
