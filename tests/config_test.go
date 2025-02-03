// File: tests/config_test.go
package tests

import (
	"fmt"
	"github.com/manasrb21/blitzconf/blitzconf"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYAMLConfig(t *testing.T) {
	filePath := "../examples/config.yaml"
	fmt.Printf("🔍 Attempting to load YAML file: %s\n", filePath)
	cfg, err := blitzconf.Load(filePath)
	assert.NoError(t, err, "Loading YAML config should not fail")

	assert.Equal(t, 8080, cfg.GetInt("server.port"))
	assert.Equal(t, "localhost", cfg.GetString("database.host"))
}

func TestLoadJSONConfig(t *testing.T) {
	filePath := "../examples/config.json"
	fmt.Printf("🔍 Attempting to load JSON file: %s\n", filePath)
	cfg, err := blitzconf.Load(filePath) // Assume JSON file exists
	assert.NoError(t, err, "Loading JSON config should not fail")

	assert.Equal(t, 8080, cfg.GetInt("server.port"))
	assert.Equal(t, "localhost", cfg.GetString("database.host"))
}

func TestEnvOverride(t *testing.T) {
	filePath := "../examples/config.yaml"
	fmt.Printf("🔍 Attempting to load YAML file with ENV override: %s\n", filePath)
	os.Setenv("SERVER_PORT", "9090")
	defer os.Unsetenv("SERVER_PORT")

	cfg, err := blitzconf.Load(filePath)
	assert.NoError(t, err, "Loading YAML config should not fail")

	assert.Equal(t, 9090, cfg.GetInt("server.port"))
}

func TestMissingKey(t *testing.T) {
	filePath := "../examples/config.yaml"
	fmt.Printf("🔍 Attempting to load YAML file for missing key test: %s\n", filePath)
	cfg, err := blitzconf.Load(filePath)
	assert.NoError(t, err, "Loading YAML config should not fail")

	assert.Equal(t, "", cfg.GetString("invalid.key"))
	assert.Equal(t, 0, cfg.GetInt("invalid.key"))
}
