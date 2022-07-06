package config

import (
	"log"

	"github.com/GustavoFreitas22/the_guardian/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	err      error
)

func DbConnection() {
	envToConnect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	Database, err = gorm.Open(postgres.Open(envToConnect))

	if err != nil {
		log.Println(err)
	}

	Database.AutoMigrate(&model.Data{})
}
