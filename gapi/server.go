package gapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mshero7/simplebank/db/sqlc"
	"github.com/mshero7/simplebank/pb"
	"github.com/mshero7/simplebank/token"
	"github.com/mshero7/simplebank/util"
)

// Server serves gRPC requests for out banking services
type Server struct {
	pb.UnimplementedSimpleBankServer // embeding, 구현한 RPC 이용 가능
	config                           util.Config
	store                            db.Store    // interact with the database processing API requests from client
	router                           *gin.Engine // help us send each API request to the correct handler for processing
	tokenMaker                       token.Maker
}

// NewServer creates a new HTTP Server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
