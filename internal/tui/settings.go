package tui

import (
	"fmt"
	"my-seo-tool/internal/config"

	"github.com/rivo/tview"
)

func NewSettingsForm(app *App) tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("API Key", app.settings.APIKey, 40, nil, func(text string) {
		app.settings.APIKey = text
	}).
		AddInputField("Rate Limit", fmt.Sprintf("%d", app.settings.RateLimit), 5, nil, func(text string) {
			// Convert to int if needed
		}).
		AddButton("Save", func() {
			err := config.SaveSettings("settings.yaml", app.settings)
			if err != nil {
				// handle error
			}
			app.SwitchToMainMenu()
		}).
		AddButton("Cancel", func() {
			app.SwitchToMainMenu()
		})

	form.SetBorder(true).SetTitle("Settings")
	return form
}
