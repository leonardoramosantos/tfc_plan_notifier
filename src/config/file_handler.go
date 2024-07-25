package config

import "os"

const ConfigFileName = "config.yaml"

func GetConfigFileData() []byte {
	var yaml_data []byte
	var data []byte
	var err error

	if _, err = os.Stat(ConfigFileName); err == nil {
		data, err = os.ReadFile(ConfigFileName)
	} else if _, err = os.Stat("/etc/" + ConfigFileName); err == nil {
		data, err = os.ReadFile("/etc/" + ConfigFileName)
	}

	if err == nil {
		yaml_data = data
	}

	log.Debugf("Config File value: \n%s", yaml_data)
	return yaml_data
}
