package config

import (
	"leonardoramosantos/tfc_plan_notifier/utils"
	"time"
)

type ConfigScan struct {
	OrganizationMatchExpr string
	WorkspaceMatchExpr    string
	TimeInterval          time.Duration
}

func GetConfigScan() *ConfigScan {
	var result = ConfigScan{}

	//TODO - set from yaml config
	result.OrganizationMatchExpr = ".*"
	result.WorkspaceMatchExpr = ".*"
	result.TimeInterval, _ = utils.ParseISODuration("PT12H")

	return &result
}
