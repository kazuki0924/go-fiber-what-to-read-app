package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	RDBConn *gorm.DB
)

type rdb struct{}

func NewRDB() RDB {
	return &rdb{}
}

func (*rdb) InitRDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DB_PORT"),
	)

	RDBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	// database.DBConn.AutoMigrate(&book.Book{})
	// fmt.Println("Database Migrated")
}

func (*rdb) CloseRDB() {
	sqlDB, err := RDBConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
