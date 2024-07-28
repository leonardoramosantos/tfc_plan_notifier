package config

// Structure to get configuration from YAML config file
type ConfigSlack struct {
	Token    string   `yaml:"token"`
	Channels []string `yaml:"channels"`
}
