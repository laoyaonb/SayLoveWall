package database

import (
	"SayLoveWall/app/models"

	"gorm.io/gorm"
)

func autoMigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Userdata{},
	)

	return err
}

func autoMigrateText(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Textdata{},
	)

	return err
}

func autoMigrateAdmin(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Admindata{},
	)

	return err
}
