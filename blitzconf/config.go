package blitzconf

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// ConfigLoader holds parsed configuration data
type ConfigLoader struct {
	data map[string]interface{}
}

// Load loads a YAML or JSON config file
func Load(configFile string) (*ConfigLoader, error) {
	fileExt := strings.ToLower(strings.Split(configFile, ".")[1])

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.New("Failed to read config file: " + err.Error())
	}

	loader := &ConfigLoader{data: make(map[string]interface{})}

	switch fileExt {
	case "yaml", "yml":
		if err := yaml.Unmarshal(data, &loader.data); err != nil {
			return nil, errors.New("Failed to parse YAML: " + err.Error())
		}
	case "json":
		if err := json.Unmarshal(data, &loader.data); err != nil {
			return nil, errors.New("Failed to parse JSON: " + err.Error())
		}
	default:
		return nil, errors.New("Unsupported config format: " + fileExt)
	}

	loader.overrideWithEnv()
	return loader, nil
}

// overrideWithEnv replaces config values with environment variables if available
func (c *ConfigLoader) overrideWithEnv() {
	for key := range c.data {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		if val, exists := os.LookupEnv(envKey); exists {
			c.data[key] = val
		}
	}
}

// Get retrieves a config value with type assertion
func (c *ConfigLoader) Get(key string) interface{} {
	return c.data[key]
}

// GetString retrieves a string value
func (c *ConfigLoader) GetString(key string) string {
	if val, ok := c.data[key].(string); ok {
		return val
	}
	return ""
}

// GetInt retrieves an integer value
func (c *ConfigLoader) GetInt(key string) int {
	if val, ok := c.data[key].(int); ok {
		return val
	}
	return 0
}
