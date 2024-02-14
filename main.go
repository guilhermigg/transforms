package main

import (
	"transforms/config"
	"transforms/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Error("config initialization error %s", err)
		return
	}

	router.Initialize()
}
