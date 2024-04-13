package main

import (
	"log"
	"os"

	"github.com/PhuPhuoc/hrm-v1/api"
	"github.com/PhuPhuoc/hrm-v1/db"
	"github.com/joho/godotenv"
)

func main() {
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatal("(From main) Error loading .env file:", err_env)
	}
	db, err_db := db.NewPostgresStore()
	if err_db != nil {
		log.Fatal("(From main) cannot connect to db: ", err_db)
	}
	server := api.NewServer(":"+os.Getenv("PORT"), db)
	if err_sv := server.Run(); err_sv != nil {
		log.Fatal("(From main) Cannot run the server: ", err_sv)
	}
}
