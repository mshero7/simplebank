package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mshero7/simplebank/db/sqlc"
)

// Server serves HTTP requests for out banking services
type Server struct {
	store  db.Store    // interact with the database processing API requests from client
	router *gin.Engine // help us send each API request to the correct handler for processing
}

// NewServer creates a new HTTP Server and setup routing
func NewSever(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
