package config

import (
	"fmt"
	"inventory-app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	username := "root"
	password := ""
	hostname := "127.0.0.1:3306"
	dbname := "inventory_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, hostname, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})

	DB = database
}
