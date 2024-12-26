package config

import (
	"fmt"
	"go-api/internal/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() error {
	dns := "root:admin@tcp(localhost:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

  // auto migrate
  DB.AutoMigrate(&entity.Product{})

	fmt.Println("Database connected")
	return nil
}
