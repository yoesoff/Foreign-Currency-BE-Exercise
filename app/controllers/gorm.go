package controllers

import (
	"Foreign-Currency-BE-Exercise/app/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Panicf("Failed PostgreSQL gorm.Open: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Exchange{})
	DB = db
}
