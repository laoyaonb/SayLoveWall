package database

import (
	"SayLoveWall/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error{
	err := db.AutoMigrate(
		&models.Userdata{},
	)

	return err
}