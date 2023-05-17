package users

import (
	"context"
	"log"

	"github.com/diazharizky/use-case-inventory-management/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv server) Create(ctx context.Context, newUser *pb.User) (*emptypb.Empty, error) {
	var stat *status.Status

	if err := srv.appCtx.UserRepository.Create(newUser); err != nil {
		log.Printf("Error unable to create user record: %s\n", err.Error())

		stat = status.New(codes.Internal, "Internal server error")
		return nil, stat.Err()
	}

	return nil, nil
}
