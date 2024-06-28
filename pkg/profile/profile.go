package profile

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name     string                     `json:"name"`
	Context  string                     `json:"context"`
	Services map[string][]ConfigService `json:"services"`
}

type ConfigService struct {
	Name string `json:"name"`
	// IP    string `json:"ip"`
	// Alias string `json:"alias,omitempty"`
}

func LoadConfig(filepath string) (*Config, error) {
	cfg := &Config{}

	dat, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, fmt.Errorf("profile: %v", err)
	}

	err = yaml.Unmarshal(dat, cfg)
	if err != nil {
		return cfg, fmt.Errorf("profile: %v", err)
	}

	return cfg, nil
}

func (c *Config) GetNamespaces() []string {
	namespaces := []string{}
	for namespace := range c.Services {
		namespaces = append(namespaces, namespace)
	}
	return namespaces
}
