package repositories

import (
	"database/sql"
	"time"

	"github.com/diazharizky/use-case-inventory-management/pb"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userRepository {
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

func (r userRepository) Create(newUser *pb.User) error {
	sqlStatement := `INSERT INTO users (username, full_name, email, created_at) VALUES ($1, $2, $3, $4)`

	now := time.Now()
	_, err := r.db.Exec(
		sqlStatement,
		newUser.Username,
		newUser.FullName,
		newUser.Email,
		now,
	)

	return err
}
