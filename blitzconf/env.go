package blitzconf

import (
	"os"
	"strings"
)

// OverrideWithEnv replaces config values with environment variables if available
func OverrideWithEnv(data map[string]interface{}) {
	for key := range data {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		if val, exists := os.LookupEnv(envKey); exists {
			data[key] = val
		}
	}
}
