package tui

import (
	"github.com/rivo/tview"
)

func NewMainMenu(app *App) tview.Primitive {
	menu := tview.NewList().ShowSecondaryText(false)
	menu.AddItem("Generate CSV from Templates", "", '1', func() {
		ShowGenerateCSVFlow(app)
	})
	menu.AddItem("Browse Generated CSV Files", "", '2', func() {
		// TODO: Implement CSV browsing UI.
		modal := tview.NewModal().
			SetText("CSV Browsing not implemented yet.").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SwitchToMainMenu()
			})
		app.tviewApp.SetRoot(modal, true)
	})
	menu.AddItem("Browse Results", "", '3', func() {
		modal := tview.NewModal().
			SetText("Results Browsing not implemented yet.").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SwitchToMainMenu()
			})
		app.tviewApp.SetRoot(modal, true)
	})
	menu.AddItem("Visualize Data", "", '4', func() {
		modal := tview.NewModal().
			SetText("Data Visualization not implemented yet.").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SwitchToMainMenu()
			})
		app.tviewApp.SetRoot(modal, true)
	})
	menu.AddItem("Settings", "", '5', func() {
		form := NewSettingsForm(app)
		app.tviewApp.SetRoot(form, true)
	})
	menu.AddItem("View Logs", "", '6', func() {
		// For demonstration, assume log file is at logs/app.log.
		logsView := NewLogsView(app, "logs/app.log")
		app.tviewApp.SetRoot(logsView, true)
	})
	menu.AddItem("Help", "", '7', func() {
		modal := tview.NewModal().
			SetText("Help:\n1: Generate CSV\n2: Browse CSV\n3: Browse Results\n4: Visualize Data\n5: Settings\n6: Logs\n7: Help\n8: Switch Workspace\n9: Exit").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SwitchToMainMenu()
			})
		app.tviewApp.SetRoot(modal, true)
	})
	menu.AddItem("Switch Workspace", "", '8', func() {
		wsSelection := NewWorkspaceSelection(app)
		app.tviewApp.SetRoot(wsSelection, true)
	})
	menu.AddItem("Exit", "", '9', func() {
		app.tviewApp.Stop()
	})
	menu.SetBorder(true).SetTitle("Main Menu")
	return menu
}
