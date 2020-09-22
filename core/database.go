package core

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Initialize(database *gorm.DB) error {
	db = database
	db.AutoMigrate(&Dong{})
	return nil
}
