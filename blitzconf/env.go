package blitzconf

import (
	"os"
	"strconv"
	"strings"
)

// OverrideWithEnv replaces config values with environment variables if available
func OverrideWithEnv(data map[string]interface{}) {
	overrideNestedKeys("", data)
}

// Recursively traverse and override nested keys with ENV values
func overrideNestedKeys(prefix string, data map[string]interface{}) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		envKey := strings.ToUpper(strings.ReplaceAll(fullKey, ".", "_"))

		// Fetch ENV variable
		envValue, exists := os.LookupEnv(envKey)
		//fmt.Printf("🔍 Checking ENV: %s = %s (Exists: %v)\n", envKey, envValue, exists)

		if exists {
			// Convert numeric values
			if intValue, err := strconv.Atoi(envValue); err == nil {
				data[key] = intValue
			} else {
				data[key] = envValue
			}
		}

		// If the value is a nested map, recurse into it
		if nestedMap, ok := value.(map[string]interface{}); ok {
			overrideNestedKeys(fullKey, nestedMap)
		}
	}
}
