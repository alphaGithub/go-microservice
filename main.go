package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hello/config"
	"github.com/hello/routes"
)

func main() {
	config.LoadConfig()
	fmt.Println("config ->", config.Config, config.Config.Port)
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(fmt.Sprintf(":%d", config.Config.Port))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>", config.Config.Port)
}
