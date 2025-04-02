package main

import (
	"gin-api-project/config"
	"gin-api-project/routes"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
