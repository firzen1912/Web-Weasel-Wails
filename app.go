package main

import (
	"Web-Weasel/backend"
	"context"
	"fmt"

	"github.com/wailsapp/wails/runtime"
)

// App struct
type App struct {
	ctx context.Context
	runtime.Runtime
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	backend.Run()
}

func (a *App) OpenLink(url string) error {
	err := a.Browser.OpenURL(url)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
