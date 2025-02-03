package main

import (
	"fmt"
	"github.com/manasrb21/blitzconf/blitzconf"
	"log"
)

func main() {
	// Load configuration
	cfg, err := blitzconf.Load("examples/config.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Fetch values directly using dot notation
	port := cfg.GetInt("server.port")
	dbHost := cfg.GetString("database.host")

	fmt.Printf("🚀 Server running on port %d\n", port)
	fmt.Printf("🗄️ Database Host: %s\n", dbHost)
}
