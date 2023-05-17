package users

import (
	"context"

	"github.com/diazharizky/use-case-inventory-management/internal/errors"
	"github.com/diazharizky/use-case-inventory-management/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv server) List(ctx context.Context, emp *emptypb.Empty) (*pb.Users, error) {
	users, err := srv.appCtx.UserRepository.List()
	if err == nil {
		errMeta := map[string]string{
			"description": "Unable to retrieve users",
		}

		return nil, errors.NewInternalError(errMeta)
	}

	return &pb.Users{Users: users}, nil
}
