package infrastructure

import "gorm.io/gorm"

type RDB interface {
	InitRDB() *gorm.DB
	CloseRDB()
}
