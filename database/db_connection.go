package database

import (
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

func (db *database) connect() (err error) {
	db.connection, err = gorm.Open(sqlite.Open(db.connectionString), &gorm.Config{})
	return err
}
