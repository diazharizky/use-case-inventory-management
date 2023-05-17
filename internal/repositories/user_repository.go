package repositories

import (
	"github.com/diazharizky/use-case-inventory-management/pb"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (userRepository) List() ([]*pb.User, error) {
	users := []*pb.User{
		{
			Username: "aominedaiki",
			FullName: "Aomine Daiki",
			Email:    "daiki.aomine@knb.com",
		},
		{
			Username: "kiseryota",
			FullName: "Kise Ryota",
			Email:    "ryota.kise@knb.com",
		},
	}

	return users, nil
}

func (r userRepository) Create(newUser pb.User) error {
	return r.db.Create(newUser).Error
}
