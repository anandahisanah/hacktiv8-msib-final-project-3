package database

import (
	"hacktiv8-msib-final-project-3/config"
	// "hacktiv8-msib-final-project-3/entity"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
