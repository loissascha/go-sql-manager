package app

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql-manager/internal/databases"
)

type App struct {
	ctx                context.Context
	activeDb           databases.Database
	activeDbConnection *sql.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.activeDb = &databases.MySQL{}
	err := a.activeDb.SetConnectionString("root:root@tcp(127.0.0.1:30306)/")
	if err != nil {
		panic(err)
	}
	a.activeDbConnection, err = a.activeDb.Connect()
	if err != nil {
		panic(err)
	}
	// TODO: dynamically change active connection based on what the user has selected
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListDbTables() []string {
	list, err := a.activeDb.ListDatabases(a.activeDbConnection)
	if err != nil {
		panic(err)
	}
	return list
}
