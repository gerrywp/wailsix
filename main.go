package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	ncn := NewNcn()
	// Create application menu
	mainMenu := menu.NewMenu()
	viewMenu := mainMenu.AddSubmenu("NCN管理")
	viewMenu.AddText("首页", keys.CmdOrCtrl("1"), func(_ *menu.CallbackData) {
		app.navigateTo("/")
	})
	viewMenu.AddText("退出", keys.CmdOrCtrl("2"), func(_ *menu.CallbackData) {
		app.navigateTo("/login")
	})
	viewMenu.AddText("关于", nil, func(_ *menu.CallbackData) {
		app.navigateTo("/about")
	})
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wailsix",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             mainMenu,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			ncn.startup(ctx)
		},
		Bind: []interface{}{
			app,
			ncn,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
