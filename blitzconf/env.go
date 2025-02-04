package blitzconf

import (
	"os"
	"strings"
	"sync"
)

var envCache sync.Map // Store environment variables

// OverrideWithEnv replaces config values with cached environment variables
func OverrideWithEnv(data map[string]interface{}) {
	for key := range data {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

		// Check if already cached
		if cachedVal, ok := envCache.Load(envKey); ok {
			data[key] = cachedVal
			continue
		}

		// Fetch from os.Getenv and cache
		if envValue, exists := os.LookupEnv(envKey); exists {
			data[key] = envValue
			envCache.Store(envKey, envValue) // Store for faster lookups
		}
	}
}
