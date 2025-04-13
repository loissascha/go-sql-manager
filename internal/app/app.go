package app

import (
	"context"
	"database/sql"
	"go-sql-manager/internal/configs"
	"go-sql-manager/internal/databases"
	"log"
)

type App struct {
	ctx                context.Context
	activeDb           databases.Database
	activeDbConnection *sql.DB
	databaseConfig     configs.Database
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.databaseConfig = *configs.NewDatabase()

	// a.selectMysqldb()
	// a.selectPostgresqldb()
	// c, err := a.activeDb.Connect()
	// if err != nil {
	// 	panic(err)
	// }
	// a.activeDbConnection = c
	// TODO: dynamically change active connection based on what the user has selected
}

func (a *App) selectMysqldb() {
	a.activeDb = &databases.MySQL{}
	err := a.activeDb.SetConnectionString("root:root@tcp(127.0.0.1:30306)/")
	if err != nil {
		panic(err)
	}
}

func (a *App) selectPostgresqldb() {
	a.activeDb = &databases.PostgreSQL{}
	err := a.activeDb.SetConnectionString("postgres://postgres:root@localhost:5432/")
	if err != nil {
		panic(err)
	}
}

func (a *App) ListDbTables() []string {
	if a.activeDb == nil {
		return []string{}
	}
	list, err := a.activeDb.ListDatabases(a.activeDbConnection)
	if err != nil {
		panic(err)
	}
	return list
}

func (a *App) ListTables(dbName string) []string {
	if a.activeDb == nil {
		return []string{}
	}
	list, err := a.activeDb.ListTables(a.activeDbConnection, dbName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return list
}

func (a *App) GetDatabaseConfigs() []configs.DatabaseConfig {
	return a.databaseConfig.GetDatabaseConfigs()
}
