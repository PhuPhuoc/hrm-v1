package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("(From migrate) Error loading .env file:", err)
		return
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	m, err := migrate.New(
		"file://migrate/migrations",
		connStr)
	if err != nil {
		log.Fatal("err new migrate: ", err)
	}
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("err migrate up: ", err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("err migrate down: ", err)
		}
	}

	if cmd == "down-one" {
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal("err rolling back one step: ", err)
		}
	}

	m.Close()
	/*
		migrate create -ext sql -dir migrate/migrations -seq add_account_table
		go run migrate/main.go up
		go run migrate/main.go down
		go run migrate/main.go down-one
	*/
}
