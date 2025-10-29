package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hello/config"
	"github.com/hello/databases"
	"github.com/hello/routes"
)

func initServer() {
	config.LoadConfig()
	databases.InitDb()
}

func main() {
	initServer()

	server := gin.Default()

	// register routes
	routes.RegisterRoutes(server)
	err := server.Run(fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		fmt.Println("❌ server startup failed >>>>>", err)
	}
	fmt.Println("✅ server started at port >>>>>", config.Config.Port)
}
