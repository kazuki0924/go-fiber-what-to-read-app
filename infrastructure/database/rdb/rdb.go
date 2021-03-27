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

func (*rdb) InitRDB() *gorm.DB {
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
		if err != nil {
			panic("Failed to connect to database")
		}
		return nil
	}
	fmt.Println("Database connection successfully opened")

	return RDBConn
}

func (*rdb) CloseRDB() {
	sqlDB, err := RDBConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}

func (*rdb) SetupMigrations(db *gorm.DB, dst ...interface{}) {
	db.AutoMigrate(dst...)

	fmt.Println("Database Migrated")
}
