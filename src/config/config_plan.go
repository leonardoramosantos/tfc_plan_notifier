package config

import (
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

// Structure to get configuration from YAML config file
type ConfigPlan struct {
	TFCToken string       `yaml:"tfc-token"`
	Scans    []ConfigScan `yaml:"scans"`
}

func (s *ConfigPlan) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(s)

	type plain ConfigPlan
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	return nil
}

func GetConfigPlan(cfg string) *ConfigPlan {
	var result ConfigPlan

	yaml.Unmarshal(GetConfigFileData(), &result)

	log.Debugf("Plans: %s", result)

	return &result
}
