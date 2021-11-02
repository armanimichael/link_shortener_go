package database

import (
	"github.com/armanimichael/link_shortener_go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sync"
)

const defaultConnection string = "./linkshortener.db"

// DB connection singleton instance
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
	if err := connection.AutoMigrate(&models.Link{}); err != nil {
		log.Fatal(err)
	}
}

func connectAndMigrate() {
	connect()
	migrate()
}

// GetDB returns the singleton DB instance
func GetDB() *gorm.DB {
	once.Do(connectAndMigrate)
	return connection
}
