package main

import (
	"github.com/gustavohroos/gopportunities/config"
	"github.com/gustavohroos/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Init config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	logger.Info("Initializing router...")
	// Init router
	router.Initialize()
}
