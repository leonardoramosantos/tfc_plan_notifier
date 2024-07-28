package config

import (
	"github.com/creasty/defaults"
)

// Structure to get configuration from YAML config file
type ConfigScan struct {
	OrganizationMatchExpr string        `default:".*" yaml:"organization"`
	WorkspaceMatchExpr    string        `default:".*" yaml:"workspace"`
	TimeInterval          string        `default:"PT12H" yaml:"interval"`
	SlackNotifications    []ConfigSlack `yaml:"slack-notifications"`
}

func (s *ConfigScan) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(s)

	type plain ConfigScan
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	return nil
}
