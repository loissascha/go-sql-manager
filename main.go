package main

import (
	"database/sql"
	"embed"
	"fmt"
	"go-sql-manager/internal/app"
	"go-sql-manager/internal/databases"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	fmt.Println("Starting go-sql-manager...")

	db := &databases.MySQL{}
	err := db.SetConnectionString("root:root@tcp(127.0.0.1:30306)/")
	if err != nil {
		panic(err)
	}
	dbConn := connectDb(db)

	dblist, err := db.ListDatabases(dbConn)
	if err != nil {
		panic(err)
	}
	fmt.Println(dblist)

	// Create an instance of the app structure
	a := app.NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "go-sql-manager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func connectDb(d databases.Database) *sql.DB {
	db, err := d.Connect()
	if err != nil {
		panic(err)
	}
	return db
}
