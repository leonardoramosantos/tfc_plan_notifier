package config

type ConfigSlack struct {
	Token    string   `yaml:"token"`
	Channels []string `yaml:"channels"`
}
