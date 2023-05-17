package db

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	Connect() (*gorm.DB, error)
}
