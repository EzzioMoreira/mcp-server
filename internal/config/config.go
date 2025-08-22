package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Main Config struct that hold all configuration sections using struct tags to map yaml keys
type Config struct {
	Server ServerConfig `yaml:"server"`
	K8s    K8sConfig    `yaml:"kubernetes"`
	Log    LogConfig    `yaml:"logging"`
}

// Define MCP Server metadata like name, version that will be send during MCP handshake
type ServerConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

// K8s Connection kubeconfig file
type K8sConfig struct {
	ConfigPath string   `yaml:"configPath"`
	Context    string   `yaml:"context"`
	Namespaces []string `yaml:"namespaces"`
}

// Logging behavior levels
type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// Function that return configuration
func Load() (*Config, error) {
	// Default configuration values
	cfg := &Config{
		Server: ServerConfig{
			Name:        "k8s-mcp-server",
			Version:     "1.0.0",
			Description: "Kubernetes MCP Server for IA-powered cluster management",
		},
		// Sets config default file localization ~/.kube/config
		K8s: K8sConfig{
			ConfigPath: filepath.Join(os.Getenv("HOME"), ".kube", "config"),
			Namespaces: []string{"default"},
		},
		Log: LogConfig{
			Level:  "info",
			Format: "json",
		},
	}

	// Override default config if CONFIG_FILE is set
	if ConfigFile := os.Getenv("CONFIG_FILE"); ConfigFile != "" {
		data, err := os.ReadFile(ConfigFile)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}
