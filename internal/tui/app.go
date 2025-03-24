package tui

import (
	"log"

	"my-seo-tool/internal/config"

	"github.com/rivo/tview"
)

// App holds the TUI application state.
type App struct {
	tviewApp         *tview.Application
	logger           *log.Logger
	settings         *config.Settings
	workspaces       *config.WorkspacesConfig
	currentWorkspace config.WorkspaceEntry // <-- Add this field
}

// NewApp initializes and returns a new App instance.
func NewApp(settings *config.Settings, workspaces *config.WorkspacesConfig, logger *log.Logger) *App {
	return &App{
		tviewApp:   tview.NewApplication(),
		logger:     logger,
		settings:   settings,
		workspaces: workspaces,
		// currentWorkspace will be set when the user selects a workspace.
	}
}

// SwitchToMainMenu sets the main menu as the root view.
func (a *App) SwitchToMainMenu() {
	mainMenu := NewMainMenu(a)
	a.tviewApp.SetRoot(mainMenu, true)
}

// showErrorModal is a helper to display an error message.
func (a *App) showErrorModal(msg string) {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.SwitchToMainMenu()
		})
	a.tviewApp.SetRoot(modal, true)
}

// Run sets up the initial UI (e.g., workspace selection) and starts the TUI event loop.
func (a *App) Run() error {
	// For example, display the workspace selection screen:
	wsSelection := NewWorkspaceSelection(a)
	a.tviewApp.SetRoot(wsSelection, true)
	return a.tviewApp.Run()
}
