package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./server.sql")
	if err != nil {
		fmt.Println("NewDB err", err)
	}
	db.LogMode(true)
	db.AutoMigrate(&TODOModel{})

	return db
}
