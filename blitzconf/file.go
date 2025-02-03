package blitzconf

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// ReadConfigFile reads a configuration file and returns parsed data
func ReadConfigFile(configFile string) (map[string]interface{}, error) {
	fileExt := strings.ToLower(strings.Split(configFile, ".")[1])

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.New("❌ Failed to read config file: " + err.Error())
	}

	var parsedData map[string]interface{}

	switch fileExt {
	case "yaml", "yml":
		if err := yaml.Unmarshal(data, &parsedData); err != nil {
			return nil, errors.New("❌ Failed to parse YAML: " + err.Error())
		}
	case "json":
		if err := json.Unmarshal(data, &parsedData); err != nil {
			return nil, errors.New("❌ Failed to parse JSON: " + err.Error())
		}
	default:
		return nil, errors.New("❌ Unsupported config format: " + fileExt)
	}

	return parsedData, nil
}
