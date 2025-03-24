package config

// Settings, ClustersConfig, Template, and TemplatesConfig definitions.
type Settings struct {
	APIKey            string            `yaml:"apiKey"`
	RateLimit         int               `yaml:"rateLimit"`
	DefaultOutputDir  string            `yaml:"defaultOutputDir"`
	DefaultResultsDir string            `yaml:"defaultResultsDir"`
	KeyBindings       map[string]string `yaml:"keyBindings"`
}

type ClustersConfig struct {
	Clusters map[string][]string `yaml:"clusters"`
}

type Template struct {
	Name     string   `yaml:"name"`
	Sequence []string `yaml:"sequence"`
}

type TemplatesConfig struct {
	Templates []Template `yaml:"templates"`
}
