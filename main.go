package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hello/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":5890")
}
