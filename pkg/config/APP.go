package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	// db, err := gorm.Open(typeof db connecting to --> "mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", /**/)
		if err != nil {
			panic("failed to connect database")
		}
	db = d
}

func GetDb() *gorm.DB {
	return db
}