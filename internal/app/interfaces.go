package app

import "github.com/diazharizky/use-case-inventory-management/pb"

type IUserRepository interface {
	List() ([]*pb.User, error)
	Create(pb.User) error
}
