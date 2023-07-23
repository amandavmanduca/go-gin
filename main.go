package main

import (
	"github.com/amandavmanduca/api-go-gin/database"
	"github.com/amandavmanduca/api-go-gin/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
