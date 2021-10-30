package main

import (
	"github.com/armanimichael/link_shortener_go/database"
	"github.com/armanimichael/link_shortener_go/routes"
)

const connectionStr string = "linkshortener.db"

func main() {
	db := setupDB()
	router := routes.NewRouter(db)
	router.Run()
}

func setupDB() *database.Database {
	db := database.NewDatabase(connectionStr)
	db.Connect()
	db.Migrate()
	return db
}
