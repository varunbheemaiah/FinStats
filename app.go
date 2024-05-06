package main

import (
	"context"
	"embed"
	"encoding/base64"
	"fmt"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed images/*
var images embed.FS

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SaveFile() string {
	selection, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Select File",
		DefaultFilename: "Earnings Calendar.xlsx",
		Filters: []runtime.FileFilter{
			{DisplayName: "Excel '97-2004 Workbooks (*.xls)", Pattern: "*.xls"},
			{DisplayName: "Excel Workbooks (*.xlsx)", Pattern: "*.xlsx"},
			{DisplayName: "Excel Binary Workbooks (*.xlsb)", Pattern: "*.xlsb"},
			{DisplayName: "Numbers Spreadsheets (*.numbers)", Pattern: "*.numbers"},
		},
	})
	if err != nil {
		_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Selection Error",
			Message: err.Error(),
		})
		return ""
	}
	return selection
}

func (a *App) WriteFile(b64 string, path string) {
	buf, _ := base64.StdEncoding.DecodeString(b64)
	_ = os.WriteFile(path, buf, 0644)
}

func (a *App) ShowError(title string, message string) {
	_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   title,
		Message: message,
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
