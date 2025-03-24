package tui

import (
	"github.com/rivo/tview"
)

func NewWorkspaceSelection(app *App) tview.Primitive {
	list := tview.NewList().ShowSecondaryText(false)
	// Add each workspace from the global config.
	for _, ws := range app.workspaces.Workspaces {
		// Capture the workspace name in the closure.
		wsName := ws.Name
		list.AddItem(wsName, "", 0, func() {
			// Here, load workspace-specific clusters/templates if needed.
			app.logger.Printf("Selected workspace: %s", wsName)
			// Switch to the workspace main menu.
			app.SwitchToMainMenu()
		})
	}
	list.AddItem("Exit", "", 'q', func() {
		app.tviewApp.Stop()
	})
	list.SetBorder(true).SetTitle("Select Workspace")
	return list
}
