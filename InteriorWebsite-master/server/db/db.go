package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dataSourceName := "root:password@tcp(localhost:3306)/interior?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dataSourceName))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	log.Default().Println("Closing database connection")
	sqlDB.Close()
	log.Default().Println("Connection closed")
}
