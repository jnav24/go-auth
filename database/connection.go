package database

import (
	"../models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
	// change username, password, and db credentials
	connection, err := gorm.Open(mysql.Open("username:password@/database"), &gorm.Config({}}))

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}