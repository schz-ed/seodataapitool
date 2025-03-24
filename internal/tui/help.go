package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewHelpScreen(app *App) tview.Primitive {
	helpText := `Help & Shortcuts:

1: Generate CSV from selected templates
2: Browse Generated CSV Files
3: Browse Results
4: Visualize Data
5: Settings
6: View Logs
7: Help
8: Switch Workspace
9: Exit

Press any key to return.`
	textView := tview.NewTextView().SetText(helpText)
	textView.SetBorder(true).SetTitle("Help")
	textView.SetDoneFunc(func(key tcell.Key) {
		app.SwitchToMainMenu()
	})
	return textView
}
