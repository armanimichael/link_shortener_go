package database

import (
	"fmt"
	"github.com/armanimichael/link_shortener_go/database/models"
	"testing"
)

func getDB() *database {
	dir := fmt.Sprintf("../%s", defaultConnection)
	db := newDatabase(dir)
	return db
}

func TestDatabase_Connect(t *testing.T) {
	db := getDB()
	err := db.connect()

	if err != nil {
		t.Errorf("failed to connect database")
	}
}

func TestDatabase_Migrate(t *testing.T) {
	db := getDB()
	db.connect()
	db.migrate()

	var links []models.Link
	result := db.connection.Find(&links)
	if result.Error != nil {
		t.Errorf("table link does not exist")
	}
}
