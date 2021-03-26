package infrastructure

import (
	"fmt"
	"log"
	"os"

	model "github.com/kazuki0924/go-what-to-read-app/domain/model"
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

func (*rdb) InitRDB() (*gorm.DB, error) {
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
		return nil, err
	}
	fmt.Println("Database connection successfully opened")

	RDBConn.AutoMigrate(&model.Book{})
	fmt.Println("Database Migrated")

	return RDBConn, nil
}

func (*rdb) CloseRDB() {
	sqlDB, err := RDBConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
