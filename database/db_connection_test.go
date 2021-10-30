package database

import (
	"fmt"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	conn := fmt.Sprintf("../%s", defaultConnection)
	db := NewDatabase(conn)
	err := db.connect()

	if err != nil {
		t.Errorf("failed to connect database")
	}
}
