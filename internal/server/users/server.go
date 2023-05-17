package users

import (
	"github.com/diazharizky/use-case-inventory-management/internal/app"
	"github.com/diazharizky/use-case-inventory-management/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer

	appCtx *app.Context
}

func NewServer(appCtx *app.Context) *server {
	return &server{
		appCtx: appCtx,
	}
}
