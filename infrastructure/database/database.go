package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DB_PORT"),
	)

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	// database.DBConn.AutoMigrate(&book.Book{})
	// fmt.Println("Database Migrated")
}

func CloseDB() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
