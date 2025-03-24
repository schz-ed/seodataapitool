// config.go or types.go (same package config)
package config

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
