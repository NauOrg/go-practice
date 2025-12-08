package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Config struct {
	FeatureEnabled bool
	MaxUsers       int
	Theme          string
}

var config atomic.Value

func serveUser(id int) {
	for {
		cfg := config.Load().(Config) // no lock!

		if cfg.FeatureEnabled {
			fmt.Printf("[User %d] NEW feature | Theme: %s | MaxUsers: %d\n", id, cfg.Theme, cfg.MaxUsers)
		} else {
			fmt.Printf("[User %d] OLD feature | Theme: %s | MaxUsers: %d\n", id, cfg.Theme, cfg.MaxUsers)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func reloadConfig() {
	time.Sleep(2 * time.Second) // simulate remote download
	newCfg := Config{
		FeatureEnabled: true,
		MaxUsers:       500,
		Theme:          "dark",
	}
	fmt.Println("\n🔄 CONFIG UPDATED LIVE\n")
	config.Store(newCfg) // atomic swap
}

func main() {
	// Initial config
	config.Store(Config{
		FeatureEnabled: false,
		MaxUsers:       100,
		Theme:          "light",
	})

	// Start "users" reading config
	for i := 1; i <= 3; i++ {
		go serveUser(i)
	}

	// Update config in background
	go reloadConfig()

	select {} // run forever
}
