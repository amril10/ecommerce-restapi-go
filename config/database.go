package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {

	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	DB = db

	err = db.AutoMigrate(
	// &model.Role{},
	// &model.User{},
	// &model.Category{},
	// &model.Product{},
	// &model.Cart{},
	)

	if err != nil {
		log.Fatalf("Automigrate error %v", err)
	}

}
