package tui

import (
	"fmt"
	"os"
	"path/filepath"

	"my-seo-tool/internal/config"
	"my-seo-tool/internal/generator"

	"github.com/rivo/tview"
)

// ShowGenerateCSVFlow loads clusters and templates, presents a selection menu,
// and generates CSV files for the selected templates.
func ShowGenerateCSVFlow(app *App) {
	// Assume currentWorkspace is set in your App struct.
	ws := app.currentWorkspace

	// Load clusters from the workspace's clusters file.
	clustersCfg, err := config.LoadClusters(ws.ClustersFile)
	if err != nil {
		app.showErrorModal(fmt.Sprintf("Error loading clusters: %v", err))
		return
	}

	// Load templates from the workspace's templates file.
	templatesCfg, err := config.LoadTemplates(ws.TemplatesFile)
	if err != nil {
		app.showErrorModal(fmt.Sprintf("Error loading templates: %v", err))
		return
	}

	// Create a selection list for templates.
	list := tview.NewList().ShowSecondaryText(false)
	// Use a map to track which templates are selected (by index).
	selected := make(map[int]bool)

	// List each template.
	for i, tmpl := range templatesCfg.Templates {
		idx := i // capture loop variable
		list.AddItem(tmpl.Name, "Press Enter to toggle selection", 0, func() {
			if selected[idx] {
				delete(selected, idx)
			} else {
				selected[idx] = true
			}
			// (Optional) You might update the UI to reflect selection changes.
		})
	}

	// Option to generate CSV files from the selected templates.
	list.AddItem("Generate CSV for selected templates", "", 0, func() {
		outputDir := filepath.Join(app.settings.DefaultOutputDir, ws.Name)
		// Ensure the output directory exists.
		os.MkdirAll(outputDir, 0755)

		// Iterate over the selected templates and generate CSV files.
		for i, tmpl := range templatesCfg.Templates {
			if selected[i] {
				err := generator.GenerateAndSave(clustersCfg, tmpl, outputDir)
				if err != nil {
					app.logger.Printf("Error generating CSV for %s: %v", tmpl.Name, err)
				} else {
					app.logger.Printf("Generated CSV for %s", tmpl.Name)
				}
			}
		}

		// Show a success modal and return to the main menu.
		modal := tview.NewModal().
			SetText("CSV generation complete!").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SwitchToMainMenu()
			})
		app.tviewApp.SetRoot(modal, true)
	})

	// Cancel option to return to the main menu.
	list.AddItem("Cancel", "", 0, func() {
		app.SwitchToMainMenu()
	})

	list.SetBorder(true).SetTitle("Select Templates (Enter to toggle selection)")
	app.tviewApp.SetRoot(list, true)
}
