package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DBNAME := os.Getenv("DB_DBNAME")
	DB_PORT := os.Getenv("DB_PORT")

	Connect(DB_HOST, DB_USER, DB_PASSWORD, DB_DBNAME, DB_PORT)
}

func Connect(DB_HOST, DB_USER, DB_PASSWORD, DB_DBNAME, DB_PORT string) {
	//datasource properties
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", DB_HOST, DB_USER, DB_PASSWORD, DB_DBNAME, DB_PORT)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = d
}

func GETDB() *gorm.DB {
	return DB
}
