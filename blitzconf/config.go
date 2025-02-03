package blitzconf

import (
	"fmt"
	"strings"
)

// ConfigLoader holds parsed configuration data
type ConfigLoader struct {
	data map[string]interface{}
}

// Load loads a YAML or JSON config file
func Load(configFile string) (*ConfigLoader, error) {
	parsedData, err := ReadConfigFile(configFile)
	if err != nil {
		return nil, err
	}

	OverrideWithEnv(parsedData)

	return &ConfigLoader{data: parsedData}, nil
}

// Get retrieves a config value, supporting dot-notation (e.g., "server.port")
func (c *ConfigLoader) Get(key string) interface{} {
	keys := strings.Split(key, ".")
	val := c.data

	for _, k := range keys {
		nextVal, exists := val[k]
		if !exists {
			fmt.Printf("⚠️ Key not found: %s\n", key)
			return nil
		}

		// If it's a nested map, continue navigating
		if nestedMap, ok := nextVal.(map[string]interface{}); ok {
			val = nestedMap
		} else {
			// If we reached the final value, return it
			return nextVal
		}
	}
	return val
}

// GetString retrieves a string value safely
func (c *ConfigLoader) GetString(key string) string {
	val := c.Get(key)
	if str, ok := val.(string); ok {
		return str
	}
	fmt.Printf("⚠️ Type assertion failed for key: %s\n", key)
	return ""
}

// GetInt retrieves an integer value safely, handling float64 conversions
func (c *ConfigLoader) GetInt(key string) int {
	val := c.Get(key)
	switch v := val.(type) {
	case int:
		return v
	case float64:
		return int(v) // Convert float64 to int
	default:
		fmt.Printf("⚠️ Type assertion failed for key: %s\n", key)
		return 0
	}
}
