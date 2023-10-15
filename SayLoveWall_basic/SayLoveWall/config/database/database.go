package database

import (
	"SayLoveWall/config/config/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB1 *gorm.DB //Userdata
var DB2 *gorm.DB //Textdata
var DB3 *gorm.DB //Admindata

func Init() {
	//用户信息数据库
	user1 := config.Config.GetString("MySQL.user")
	pass1 := config.Config.GetString("MySQL.password")
	host1 := config.Config.GetString("MySQL.host")
	port1 := config.Config.GetString("MySQL.port")
	DBname1 := config.Config.GetString("MySQL.DBname")

	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user1, pass1, host1, port1, DBname1)

	db1, err1 := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err1 != nil {
		panic("failed to connect database, error = " + err1.Error())
	}

	err1 = autoMigrateUser(db1)
	if err1 != nil {
		log.Fatal("Database migrate failed: ", err1)

	}
	DB1 = db1

	//文本信息数据库
	user2 := config.Config.GetString("MySQL.user")
	pass2 := config.Config.GetString("MySQL.password")
	host2 := config.Config.GetString("MySQL.host")
	port2 := config.Config.GetString("MySQL.port")
	DBname2 := config.Config.GetString("MySQL.DBname")

	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user2, pass2, host2, port2, DBname2)

	db2, err2 := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err2 != nil {
		panic("failed to connect database, error = " + err2.Error())
	}

	err2 = autoMigrateText(db2)
	if err2 != nil {
		log.Fatal("Database migrate failed: ", err2)

	}
	DB2 = db2

	//管理员信息数据库
	user3 := config.Config.GetString("MySQL.user")
	pass3 := config.Config.GetString("MySQL.password")
	host3 := config.Config.GetString("MySQL.host")
	port3 := config.Config.GetString("MySQL.port")
	DBname3 := config.Config.GetString("MySQL.DBname")

	dsn3 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user3, pass3, host3, port3, DBname3)

	db3, err3 := gorm.Open(mysql.Open(dsn3), &gorm.Config{})
	if err3 != nil {
		panic("failed to connect database, error = " + err3.Error())
	}

	err3 = autoMigrateAdmin(db3)
	if err3 != nil {
		log.Fatal("Database migrate failed: ", err3)

	}
	DB3 = db3
}
