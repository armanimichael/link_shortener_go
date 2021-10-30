package database

import (
	"github.com/armanimichael/link_shortener_go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultConnection string = "linkshortener.db"

type database struct {
	connection       *gorm.DB
	connectionString string
}

func NewDatabase(connectionString string) *database {
	return &database{
		connectionString: connectionString,
	}
}

func (db *database) Connect() (err error) {
	db.connection, err = gorm.Open(sqlite.Open(db.connectionString), &gorm.Config{})
	return err
}

func (db *database) Migrate() {
	db.connection.AutoMigrate(&models.Link{})
}
