package infrastructure

import "gorm.io/gorm"

type RDB interface {
	InitRDB() *gorm.DB
	CloseRDB()
	SetupMigrations(db *gorm.DB, dst ...interface{})
}
