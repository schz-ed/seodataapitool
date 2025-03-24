package templates

import "my-seo-tool/internal/config"

// GetTemplateByName returns a pointer to the template with the given name
// from the TemplatesConfig, and a boolean indicating whether it was found.
func GetTemplateByName(tc *config.TemplatesConfig, name string) (*config.Template, bool) {
	for _, tmpl := range tc.Templates {
		if tmpl.Name == name {
			return &tmpl, true
		}
	}
	return nil, false
}

// Additional template-related logic could be added here, such as filtering or validation.
