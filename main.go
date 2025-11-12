package main

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hello/config"
	"github.com/hello/databases"
	"github.com/hello/graph"
	"github.com/hello/middlewares"

	"github.com/hello/routes"
)

func initServer() {
	config.LoadConfig()
	databases.InitDb()
}

func main() {
	initServer()

	server := gin.Default()

	gServ := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// websocket
	server.GET("/socket", middlewares.SocketHandler)

	// GraphQL Routes
	server.POST("/query", func(c *gin.Context) {
		gServ.ServeHTTP(c.Writer, c.Request)
	})
	server.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// register routes
	routes.RegisterRoutes(server)
	err := server.Run(fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		fmt.Println("❌ server startup failed >>>>>", err)
	}
	fmt.Println("✅ server started at port >>>>>", config.Config.Port)
}
