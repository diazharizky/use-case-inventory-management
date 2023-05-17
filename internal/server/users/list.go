package users

import (
	"context"
	"log"

	"github.com/diazharizky/use-case-inventory-management/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv server) List(ctx context.Context, emp *emptypb.Empty) (*pb.ListResponse, error) {
	var stat *status.Status

	users, err := srv.appCtx.UserRepository.List()
	if err != nil {
		log.Printf("Error unable to retrieve user records: %s\n", err.Error())

		stat = status.New(codes.Internal, "Internal server error")
		return nil, stat.Err()
	}

	return &pb.ListResponse{
		Ok:    true,
		Users: users,
	}, nil
}
