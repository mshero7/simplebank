package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)
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
