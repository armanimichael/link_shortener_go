package main

import (
	router2 "github.com/armanimichael/link_shortener_go/router"
)

func main() {
	router := router2.NewRouter()
	router.Run()
}
