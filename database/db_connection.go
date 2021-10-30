package database

import (
	"github.com/armanimichael/link_shortener_go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultConnection string = "linkshortener.db"

type Database struct {
	connection       *gorm.DB
	connectionString string
}

func NewDatabase(connectionString string) *Database {
	return &Database{
		connectionString: connectionString,
	}
}

func (db *Database) Connect() (err error) {
	db.connection, err = gorm.Open(sqlite.Open(db.connectionString), &gorm.Config{})
	return err
}

func (db *Database) Migrate() {
	db.connection.AutoMigrate(&models.Link{})
}
