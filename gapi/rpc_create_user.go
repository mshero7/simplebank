package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/mshero7/simplebank/db/sqlc"
	"github.com/mshero7/simplebank/pb"
	"github.com/mshero7/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Use Getter
	hashsedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed Hash Password : %s", err)
	}

	// Use Getter
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashsedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists : %s", err)
			}
		}

		return nil, status.Errorf(codes.Internal, "Failed to Create User : %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}