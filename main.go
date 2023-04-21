package main

import (
	"fmt"
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/server"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func initDb() {
	if db.Db != nil {
		log.Println("DB has already been initialized. Skipping initialization.")
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		"5432",
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := conn.AutoMigrate(db.ModelsToMigrate()...); err != nil {
		log.Fatal(err)
	}

	db.Db = conn
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	initDb()
	server.Init()
}
