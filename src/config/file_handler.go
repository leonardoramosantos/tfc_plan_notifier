package config

import "os"

const ConfigFileName = "config.yaml"

// Method to load config file
func GetConfigFileData() []byte {
	var yaml_data []byte
	var data []byte
	var err error

	// Tests two paths to load the file
	if _, err = os.Stat(ConfigFileName); err == nil {
		data, err = os.ReadFile(ConfigFileName)
	} else if _, err = os.Stat("/etc/" + ConfigFileName); err == nil {
		data, err = os.ReadFile("/etc/" + ConfigFileName)
	}

	if err == nil {
		yaml_data = data
	} else {
		log.Errorf("Error loading config Data. Err: %s", err)
		os.Exit(-2)
	}

	log.Debugf("Config File value: \n%s", yaml_data)
	return yaml_data
}
