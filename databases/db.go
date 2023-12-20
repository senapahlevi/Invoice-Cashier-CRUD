package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func SetDB() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error", err)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbsetting := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbsetting), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &DB{
		DB: db,
	}, nil
}
