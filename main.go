package main

import (
	"embed"
	"fmt"
	"go-sql-manager/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	fmt.Println("Starting go-sql-manager...")

	// Create an instance of the app structure
	a := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-sql-manager",
		Width:  1224,
		Height: 968,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
