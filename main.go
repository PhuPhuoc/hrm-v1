package main

import (
	"log"
	"os"

	"github.com/PhuPhuoc/hrm-v1/api"
	"github.com/PhuPhuoc/hrm-v1/db"
	"github.com/joho/godotenv"
)

// @title						Human Resources Management System
// @version					1.0
// @description				Human Resource Management (HRM)
// @description				Server is a comprehensive software or hardware system designed to manage all aspects of personnel-related activities within an organization.
// @description				It encompasses employee information management, timekeeping, payroll, training and development, performance management, document management, employee interaction, security, compliance.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
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
