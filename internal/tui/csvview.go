package tui

import (
	"my-seo-tool/internal/utils"
	"path/filepath"

	"github.com/rivo/tview"
)

func NewCSVViewer(app *App, csvFile string) tview.Primitive {
	textView := tview.NewTextView().SetScrollable(true).SetWrap(true)
	records, err := utils.ReadCSV(csvFile)
	if err != nil {
		textView.SetText("Error reading CSV: " + err.Error())
	} else {
		content := ""
		for _, row := range records {
			// For demonstration, join the row fields with a space.
			content += row[0] + "\n"
		}
		textView.SetText(content)
	}
	textView.SetBorder(true).SetTitle(filepath.Base(csvFile))
	return textView
}
