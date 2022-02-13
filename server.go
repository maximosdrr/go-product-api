package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"products-api/routes"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func CreateServer() Server {
	return Server{
		port:   "3000",
		engine: gin.Default(),
	}
}

func ServerRun() {
	server := CreateServer()
	routes.GenerateRoutes(server.engine)
	log.Fatal(server.engine.Run(":" + server.port))
}
