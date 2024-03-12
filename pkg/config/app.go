package config

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	DBURL,exists := os.LookupEnv("DB_URL")
	if !exists {
		fmt.Println("DB_URL not set in .env file")
	}
	database, err := gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}