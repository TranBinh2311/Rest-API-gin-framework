package api

import (
	"github.com/example/gin_framework/initializers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// initialize the server instance
func Initial() {

	server := &Server{}
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.MigrateEnvVariables()
	server.initRouter()
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// common to handle errors Response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
