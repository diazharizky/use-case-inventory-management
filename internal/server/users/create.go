package users

import (
	"context"

	"github.com/diazharizky/use-case-inventory-management/internal/errors"
	"github.com/diazharizky/use-case-inventory-management/pb"
)

func (srv server) Create(ctx context.Context, newUser *pb.User) (*pb.User, error) {
	if err := srv.appCtx.UserRepository.Create(newUser); err != nil {
		errMeta := map[string]string{
			"description": "Unable to create record",
		}

		return nil, errors.NewInternalError(errMeta)
	}

	return newUser, nil
}
