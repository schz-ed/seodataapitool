package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadSettings(path string) (*Settings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Settings
	if err = yaml.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func SaveSettings(path string, s *Settings) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func LoadClusters(path string) (*ClustersConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cc ClustersConfig
	if err := yaml.Unmarshal(data, &cc); err != nil {
		return nil, err
	}
	return &cc, nil
}

func LoadTemplates(path string) (*TemplatesConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tc TemplatesConfig
	if err := yaml.Unmarshal(data, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}
