package main

import (
	"github.com/diazharizky/use-case-inventory-management/internal/app"
	"github.com/diazharizky/use-case-inventory-management/internal/repositories"
	"github.com/diazharizky/use-case-inventory-management/internal/server"
	"github.com/diazharizky/use-case-inventory-management/pkg/db"
)

func main() {
	appCtx := app.NewContext()

	dbConn := db.GetConnection()

	appCtx.UserRepository = repositories.NewUserRepository(dbConn)

	server.Run(appCtx)
}
