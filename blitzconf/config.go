package blitzconf

import (
	"encoding/json"
	"fmt"
	"sync"

	"golang.org/x/exp/mmap" // Import memory-mapped files for fast reads
	"gopkg.in/yaml.v3"
)

// ConfigLoader holds the loaded configuration
type ConfigLoader struct {
	data map[string]interface{}
}

// Global Cache
var (
	configCache map[string]interface{}
	once        sync.Once
)

// Load parses the config file and caches it using mmap
func Load(configFile string) (*ConfigLoader, error) {
	var err error
	once.Do(func() {
		configCache, err = readConfigFile(configFile)
		if err != nil {
			fmt.Printf("❌ Failed to parse config file: %v\n", err)
		}
	})

	if err != nil {
		return nil, err
	}

	OverrideWithEnv(configCache)
	return &ConfigLoader{data: configCache}, nil
}

// readConfigFile reads and unmarshals YAML/JSON using memory-mapped files
func readConfigFile(configFile string) (map[string]interface{}, error) {
	reader, err := mmap.Open(configFile) // Open file as memory-mapped
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to mmap config file: %w", err)
	}
	defer reader.Close()

	data := make([]byte, reader.Len())
	_, err = reader.ReadAt(data, 0) // Read the entire file
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to read mmap data: %w", err)
	}

	parsedData := make(map[string]interface{})
	switch {
	case configFile[len(configFile)-4:] == "yaml" || configFile[len(configFile)-3:] == "yml":
		if err := yaml.Unmarshal(data, &parsedData); err != nil {
			return nil, fmt.Errorf("❌ YAML parsing failed: %w", err)
		}
	case configFile[len(configFile)-4:] == "json":
		if err := json.Unmarshal(data, &parsedData); err != nil {
			return nil, fmt.Errorf("❌ JSON parsing failed: %w", err)
		}
	default:
		return nil, fmt.Errorf("❌ Unsupported config format")
	}
	return parsedData, nil
}
