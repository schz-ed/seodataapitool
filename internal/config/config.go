package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Settings holds configurable parameters.
type Settings struct {
	APIKey            string            `yaml:"apiKey"`
	RateLimit         int               `yaml:"rateLimit"`
	DefaultOutputDir  string            `yaml:"defaultOutputDir"`
	DefaultResultsDir string            `yaml:"defaultResultsDir"`
	KeyBindings       map[string]string `yaml:"keyBindings"`
}

func LoadSettings(path string) (*Settings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Settings
	err = yaml.Unmarshal(data, &s)
	if err != nil {
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
