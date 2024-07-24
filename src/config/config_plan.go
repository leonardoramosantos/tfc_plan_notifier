package config

type ConfigPlan struct {
	Plans []ConfigScan
}

func GetConfigPlan(cfg string) *ConfigPlan {
	var result = ConfigPlan{}

	//TODO - Use string
	result.Plans = append(result.Plans, *GetConfigScan())
	log.Debugf("Plans: %s", result.Plans)

	return &result
}
