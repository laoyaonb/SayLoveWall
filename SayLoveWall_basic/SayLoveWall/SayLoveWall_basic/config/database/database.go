package database

import (
	"SayLoveWall/config/config/config"
	"log"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := config.Config.GetString("MySQL.user")
	pass := config.Config.GetString("MySQL.password")
	host := config.Config.GetString("MySQL.host")
	port := config.Config.GetString("MySQL.port")
	DBname := config.Config.GetString("MySQL.DBname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, pass, host, port, DBname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error = " + err.Error())
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("Database migrate failed: ", err)

	}
	DB = db
}
