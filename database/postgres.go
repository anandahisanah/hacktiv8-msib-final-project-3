package database

import (
	"errors"
	"hacktiv8-msib-final-project-3/config"
	"hacktiv8-msib-final-project-3/entity"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func seedAdmin() {
	admin := &entity.User{
		FullName: "admin",
		Email:    "admin@hacktiv8.com",
		Password: "admin123",
		Role:     "admin",
	}
	if err := admin.HashPassword(); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := db.Create(admin).Error; err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Admin account seed success!")
}

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln(err.Error())
	}

	if db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedAdmin()
		}
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
