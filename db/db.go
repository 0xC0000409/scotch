package db

import (
	"github.com/0xC0000409/scotch/models"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ModelsToMigrate() []interface{} {
	return []interface{}{
		&models.User{},
	}
}

func Instance() *gorm.DB {
	return Db
}
