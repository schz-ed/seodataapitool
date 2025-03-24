package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type WorkspaceEntry struct {
	Name          string `yaml:"name"`
	ClustersFile  string `yaml:"clustersFile"`
	TemplatesFile string `yaml:"templatesFile"`
}

type WorkspacesConfig struct {
	Workspaces []WorkspaceEntry `yaml:"workspaces"`
}

func LoadWorkspaces(path string) (*WorkspacesConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var wc WorkspacesConfig
	err = yaml.Unmarshal(data, &wc)
	if err != nil {
		return nil, err
	}
	return &wc, nil
}
