package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:8889)/Heliumvue?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to initialize database")
	}
}

func Get() *gorm.DB {
	return db
}
