package main

import (
	router "github.com/armanimichael/link_shortener_go/router"
)

func main() {
	router := router.NewRouter()
	router.Run()
}
