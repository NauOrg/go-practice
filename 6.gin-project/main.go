package main

import (
	"gin-project/config"
	"gin-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	routes.AuthRoutes(router)
	router.Run(":8080")
}
