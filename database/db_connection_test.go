package database

import (
	"testing"
)

func TestDatabase_Connect(t *testing.T) {
	if GetDB().Error != nil {
		t.Errorf("failed to connect database")
	}
}
