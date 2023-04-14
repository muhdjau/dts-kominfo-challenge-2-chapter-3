package config

import (
	"challenge-chapter-3-sesi-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "365442"
	port     = 5432
	dbname   = "db-go-product"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	conn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Success connected to database!")

	db.Debug().AutoMigrate(models.Users{}, models.Products{}, models.Roles{})
}

func GetDB() *gorm.DB {
	return db
}
