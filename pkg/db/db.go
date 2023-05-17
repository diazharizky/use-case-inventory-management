package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/diazharizky/use-case-inventory-management/config"
	"github.com/diazharizky/use-case-inventory-management/internal/enum"
)

var (
	dbType enum.DbType
)

func init() {
	config.Global.SetDefault("db.type", "postgres")

	dbType = enum.DbType(config.Global.GetString("db.type"))
}

func GetConnection() (conn *sql.DB) {
	db, err := getDB()
	if err != nil {
		log.Fatalf("Error unable to get database connection: %v", err)
	}

	conn, err = db.Connect()
	if err != nil {
		log.Fatalf("Error unable to get database connection: %v", err)
	}

	return
}

func getDB() (db IDatabase, err error) {
	if !dbType.Validate() {
		err = fmt.Errorf("error invalid database type: %s", dbType.String())
		return
	}

	switch dbType {
	case enum.DbTypePostgres:
		db = NewPostgres()
	default:
		db = NewPostgres()
	}

	return
}
