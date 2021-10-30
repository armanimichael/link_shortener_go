package database

import (
	"fmt"
	"github.com/armanimichael/link_shortener_go/database/models"
	"testing"
)

func getDB() *Database {
	conn := fmt.Sprintf("../%s", defaultConnection)
	db := NewDatabase(conn)
	return db
}

func TestNewDatabase_Connect(t *testing.T) {
	db := getDB()
	err := db.Connect()

	if err != nil {
		t.Errorf("failed to connect Database")
	}
}

func TestDatabase_Migrate(t *testing.T) {
	db := getDB()
	db.Connect()
	db.Migrate()

	var links []models.Link
	result := db.connection.Find(&links)
	if result.Error != nil {
		t.Errorf("table link does not exist")
	}
}
