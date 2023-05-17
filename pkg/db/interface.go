package db

import (
	"database/sql"
)

type IDatabase interface {
	Connect() (*sql.DB, error)
}
