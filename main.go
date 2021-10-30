package main

import "github.com/armanimichael/link_shortener_go/routes"

func main() {
	router := routes.NewRouter()
	router.Run()
}
