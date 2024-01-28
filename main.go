package main

import (
	"github.com/jsGolden/go-opportunities/config"
	"github.com/jsGolden/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
