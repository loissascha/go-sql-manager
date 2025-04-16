package app

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql-manager/internal/configs"
	"go-sql-manager/internal/databases"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx                context.Context
	activeDb           databases.Database
	activeDbConnection *sql.DB
	databaseConfig     *configs.Database
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.databaseConfig = configs.NewDatabase()
}

func (a *App) ListDbTables() []string {
	if a.activeDb == nil {
		return []string{}
	}
	list, err := a.activeDb.ListDatabases(a.activeDbConnection)
	if err != nil {
		return []string{}
	}
	return list
}

func (a *App) ListTables(dbName string) []string {
	if a.activeDb == nil {
		return []string{}
	}
	list, err := a.activeDb.ListTables(a.activeDbConnection, dbName)
	if err != nil {
		return []string{}
	}
	return list
}

func (a *App) GetDatabaseConfigs() []configs.DatabaseConfig {
	return a.databaseConfig.GetDatabaseConfigs()
}

func (a *App) AddDatabaseConfig(host string, port string, user string, password string, engine string) {
	dbtype := configs.DATABASE_MYSQL
	if engine == "1" {
		dbtype = configs.DATABASE_POSTGRES
	}
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	cfg := configs.DatabaseConfig{
		Id:       id.String(),
		Port:     port,
		User:     user,
		Host:     host,
		Password: password,
		Type:     dbtype,
	}
	err = a.databaseConfig.AddDatabaseConfig(cfg)
	if err != nil {
		panic(err)
	}
}

func (a *App) ActivateConnection(id string) {
	dblist := a.databaseConfig.GetDatabaseConfigs()
	for _, cfg := range dblist {
		if cfg.Id == id {
			switch cfg.Type {
			case configs.DATABASE_MYSQL:
				a.activeDb = &databases.MySQL{}
				connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
				err := a.activeDb.SetConnectionString(connStr)
				if err != nil {
					panic(err)
				}
				break
			case configs.DATABASE_POSTGRES:
				a.activeDb = &databases.PostgreSQL{}
				connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
				err := a.activeDb.SetConnectionString(connStr)
				if err != nil {
					panic(err)
				}
				break
			}
			// connect the database!
			c, err := a.activeDb.Connect()
			if err != nil {
				fmt.Println("Connect failed!")
				panic(err)
			}
			a.activeDbConnection = c
			fmt.Println("Emitting Event!")
			runtime.EventsEmit(a.ctx, "ConnectionChanged")
			break
		}
	}
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
