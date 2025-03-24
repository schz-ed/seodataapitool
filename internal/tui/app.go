package tui

import (
	"log"

	"my-seo-tool/internal/config"

	"github.com/rivo/tview"
)

type App struct {
	settings   *config.Settings
	workspaces *config.WorkspacesConfig
	tviewApp   *tview.Application
	logger     *log.Logger
	// You can add fields for the current workspace, clusters, templates, etc.
}

func NewApp(settings *config.Settings, workspaces *config.WorkspacesConfig, logger *log.Logger) *App {
	return &App{
		settings:   settings,
		workspaces: workspaces,
		tviewApp:   tview.NewApplication(),
		logger:     logger,
	}
}

func (a *App) Run() error {
	// Start with the workspace selection UI.
	workspaceSelection := NewWorkspaceSelection(a)
	a.tviewApp.SetRoot(workspaceSelection, true)
	return a.tviewApp.Run()
}

func (a *App) SwitchToMainMenu() {
	mainMenu := NewMainMenu(a)
	a.tviewApp.SetRoot(mainMenu, true)
}
