package tui

import (
	"io/ioutil"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewLogsView(app *App, logFile string) tview.Primitive {
	content, err := ioutil.ReadFile(logFile)
	textView := tview.NewTextView().SetScrollable(true).SetWrap(true)
	if err != nil {
		textView.SetText("Error reading log file: " + err.Error())
	} else {
		textView.SetText(string(content))
	}
	textView.SetBorder(true).SetTitle("Logs")
	textView.SetDoneFunc(func(key tcell.Key) {
		app.SwitchToMainMenu()
	})
	return textView
}
