package database

import (
	"github.com/armanimichael/link_shortener_go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sync"
)

const defaultConnection string = "./linkshortener.db"

var connection *gorm.DB = nil
var once sync.Once

func connect() {
	var err error
	connection, err = gorm.Open(sqlite.Open(defaultConnection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func migrate() {
	connection.AutoMigrate(&models.Link{})
}

func connectAndMigrate() {
	connect()
	migrate()
}

func GetDB() *gorm.DB {
	once.Do(connectAndMigrate)
	return connection
}
