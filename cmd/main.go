package main

import (
	"my-seo-tool/internal/config"
	"my-seo-tool/internal/logger"
	"my-seo-tool/internal/tui"
)

func main() {
	// Initialize the logger (writes to logs/app.log)
	logr := logger.InitLogger("logs/app.log")

	// Load global settings from settings.yaml
	settings, err := config.LoadSettings("settings.yaml")
	if err != nil {
		logr.Fatal(err)
	}

	// Load workspaces configuration from configs/workspaces.yaml
	workspaces, err := config.LoadWorkspaces("workspaces.yaml")
	if err != nil {
		logr.Fatal(err)
	}

	// Create and run the TUI application
	app := tui.NewApp(settings, workspaces, logr)
	if err := app.Run(); err != nil {
		logr.Fatal(err)
	}
}
