package infrastructure

import (
	"fmt"
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
		fmt.Println("Error connecting to postgres")
		panic(err)
	}

	fmt.Println("Database connection successfully opened")

	return RDBConn
}

func (*rdb) CloseRDB() {
	sqlDB, err := RDBConn.DB()
	if err != nil {
		fmt.Println("Error closing postgres")
		panic(err)
	}
	sqlDB.Close()
}

func (*rdb) SetupMigrations(dst ...interface{}) {
	RDBConn.AutoMigrate(dst...)

	fmt.Println("Database Migrated")
}
