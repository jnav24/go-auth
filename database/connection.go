package database

import (
	"../models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect() {
	// change username, password, and db credentials
	connection, err := gorm.Open(mysql.Open("username:password@/database"), &gorm.Config({}}))

	if err != nil {
		panic("could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}