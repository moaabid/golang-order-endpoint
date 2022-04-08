package database

import (
	"log"
	"os"

	"githu.com/moaabid/golang-order-endpoint/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect database", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})

	Database = DbInstance{DB: db}
}
