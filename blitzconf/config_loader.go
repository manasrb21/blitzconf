package blitzconf

import (
	"strconv"
	"strings"
	"sync"
)

// Cache for faster nested key lookups
var keyCache sync.Map // map[string][]string

// Get retrieves a value, optimizing nested key lookups with caching
func (c *ConfigLoader) Get(key string) interface{} {
	// Check if we've already computed the key lookup
	if cachedKeys, ok := keyCache.Load(key); ok {
		return getNestedValue(c.data, cachedKeys.([]string))
	}

	// Compute key lookup and cache it
	keys := strings.Split(key, ".")
	keyCache.Store(key, keys)
	return getNestedValue(c.data, keys)
}

// Helper function to get nested values without re-splitting
func getNestedValue(data map[string]interface{}, keys []string) interface{} {
	for _, k := range keys {
		if nestedMap, ok := data[k].(map[string]interface{}); ok {
			data = nestedMap
		} else {
			return data[k]
		}
	}
	return nil
}

// GetInt retrieves an integer value, optimized for nested lookups
func (c *ConfigLoader) GetInt(key string) int {
	val := c.Get(key)
	if val == nil {
		return 0
	}

	switch v := val.(type) {
	case int:
		return v
	case float64:
		return int(v) // YAML sometimes parses numbers as float64
	case string:
		num, err := strconv.Atoi(v)
		if err == nil {
			return num
		}
	}
	return 0
}

// GetString retrieves a string value, optimized for nested lookups
func (c *ConfigLoader) GetString(key string) string {
	val := c.Get(key)
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	case []byte: // If parsed as byte slice, convert
		return string(v)
	default:
		return ""
	}
}

func (c *ConfigLoader) GetStringSlice(key string) []string {
	val := c.Get(key)
	if val == nil {
		return []string{}
	}

	switch v := val.(type) {
	case []interface{}: // Handle YAML lists
		var result []string
		for _, item := range v {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	case string: // Handle comma-separated values
		return strings.Split(v, ",")
	default:
		return []string{}
	}
}
